package main

import (
	"fmt"
  "strings"
  "path/filepath"
  "log"

	"kidiyoor.io/family-tree/pkg/types"
)

func main() {
  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Fatal(err)
  }

  m := types.Member{
    Name: "Devdas purushotham",
  }

  parent := "rangu/shesha_sahukar/muthakka"

  if _, err := os.Stat(dir + "/" + parent + "/" + stringsToLower(strings.ReplaceAll(m.Name, " ", "_"))); os.IsNotExist(err) {
  	fmt.Println("---")
  }
}
