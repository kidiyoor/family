package main

import (
	"encoding/json"
	"fmt"

	"kidiyoor.io/family-tree/pkg/members"
	"kidiyoor.io/family-tree/pkg/tree"
)

func main() {
	tree.Display(&members.Sankesh, "")
	tSankesh, _ := json.Marshal(members.Sankesh)
	fmt.Println(string(tSankesh))

	tree.Display(&members.Kanangi, "")
	tKanangi, _ := json.Marshal(members.Kanangi)
	fmt.Println(string(tKanangi))
}
