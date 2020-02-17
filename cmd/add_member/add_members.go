package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"kidiyoor.io/family-tree/pkg/types"
)

func main() {
	g := flag.String("gender", "", "Gender")
	p := flag.String("parent", "", "Parent")
	flag.Parse()

	if *p == "" {
		log.Fatal("Parent need to be passed")
	}

	if *g != "m" && *g != "f" {
		log.Fatal("Gender need to be either m/f")
	}

	name := strings.Join(flag.Args(), " ")
	if len(name) == 0 {
		log.Fatal("Name need to be passed")
	}

	m := types.Member{
		Name: name,
	}

	parent := *p

	gender := "Male"
	if *g == "f" {
		gender = "Female"
	}

	parentDir := os.Getenv("GOPATH") + "/src/kidiyoor.io/family-tree/pkg/members" + "/" + parent
	packageName := strings.ToLower(strings.ReplaceAll(m.Name, " ", ""))
	nameTitle := strings.Title(packageName)
	newDir := parentDir + "/" + packageName
	if _, err := os.Stat(newDir); !os.IsNotExist(err) {
		log.Fatal("Member already exists")
	}

	if err := os.Mkdir(newDir, 0755); err != nil {
		log.Fatal("Unable to create new Dir, ", err)
	}

	f, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/kidiyoor.io/family-tree/cmd/add_member/template_self.go.txt")
	if err != nil {
		os.Remove(newDir)
		log.Fatal(err)
		return
	}

	tmpl, err := template.New("newSelfFile").Parse(string(f))
	if err != nil {
		os.Remove(newDir)
		log.Fatal(err)
	}

	fw, err := os.Create(newDir + "/self.go")
	defer fw.Close()
	if err != nil {
		os.Remove(newDir)
		log.Fatal(err)
	}

	type replaceData struct {
		NameTitle string
		Name      string
		Package   string
		Gender    string
	}

	data := replaceData{
		Name:      m.Name,
		NameTitle: nameTitle,
		Package:   packageName,
		Gender:    "types." + gender,
	}
	if err := tmpl.Execute(fw, data); err != nil {
		os.Remove(newDir)
		log.Fatal("Unable to replace tmpl and create new file, ", err)
	}

	parentSelfFile, err := os.Open(parentDir + "/self.go")
	if err != nil {
		os.Remove(newDir)
		log.Fatal("Unable to open parentSelfFile, ", err)
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(parentSelfFile)
	var lines []string
	var line string
	var importAdded bool
	var childAdded bool
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		if !importAdded && line == ")\n" {
			importAdded = true
			lines = append(lines, "  \"kidiyoor.io/family-tree/pkg/members/"+parent+"/"+packageName+"\"\n")
			lines = append(lines, line)
			continue
		}

		if !childAdded && strings.HasPrefix(strings.TrimSpace(line), "Children") {
			childAdded = true

			files, err := ioutil.ReadDir(parentDir)
			if err != nil {
				log.Fatal(err)
			}
			children := ""
			for _, f := range files {
				childName := f.Name()
				if childName != "self.go" {
					children = children + "&" + strings.ReplaceAll(childName, "_", "") + "." + strings.Title(childName) + ", "
				}
			}

			lines = append(lines, "    Children: []*types.Member{"+children+"},\n")
			continue
		}

		lines = append(lines, line)
	}

	if err != io.EOF {
		os.Remove(newDir)
		log.Fatal("Failed to write to parent self.go, ", err)
	}

	parentSelfFileForWrite, err := os.Create(parentDir + "/self2.go")
	defer parentSelfFileForWrite.Close()
	if err != nil {
		os.RemoveAll(newDir)
		log.Fatal("Unable to open parentSelfFile, ", err)
	}

	datawriter := bufio.NewWriter(parentSelfFileForWrite)
	for _, data := range lines {
		_, _ = datawriter.WriteString(data)
	}
	datawriter.Flush()

	os.Remove(parentDir + "/self.go")
	os.Rename(parentDir+"/self2.go", parentDir+"/self.go")

	fmt.Println("Member added successfully")
}
