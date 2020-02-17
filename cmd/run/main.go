package main

import (
	"kidiyoor.io/family-tree/pkg/members"
	"kidiyoor.io/family-tree/pkg/tree"
)

func main() {
	tree.Display(&members.Root, "")
}
