package main

import (
	"encoding/json"
	"fmt"
        "flag"
	"kidiyoor.io/family-tree/pkg/members"
	"kidiyoor.io/family-tree/pkg/tree"
)

var (
  family string
  printJSON bool
)

func init() {
    flag.StringVar(&family, "family", "", "Name of the family")
    flag.BoolVar(&printJSON, "print-json", false, "Set true to output JSON only")
    flag.Parse()
}

func main() {
  if family == "" {
    fmt.Println(fmt.Sprintf("Error: name of the family required, possible values are %s, %s. Eg. --family=sankesh","sankesh", "kanangi"))
    return
  }
  if !printJSON {
    tree.Display(&members.Sankesh, "")
  }
  if family == "sankesh" { 
    tSankesh, _ := json.Marshal(members.Sankesh)
    fmt.Println(string(tSankesh))
  }
  if !printJSON {
    tree.Display(&members.Kanangi, "")
  }
  if family == "kanangi" {
    tKanangi, _ := json.Marshal(members.Kanangi)
    fmt.Println(string(tKanangi))
  }
}

