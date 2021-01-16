package main

import (
	"encoding/json"
	"fmt"

	"kidiyoor.io/family-tree/pkg/members"
	"kidiyoor.io/family-tree/pkg/tree"
)

func main() {
	tree.Display(&members.Root, "")
	t, _ := json.Marshal(members.Root)
	fmt.Println(string(t))
}
