package tree

import (
	"fmt"

	"kidiyoor.io/family-tree/pkg/types"
)

// Display the family tree.
func Display(m *types.Member, spaces string) {
	fmt.Println(spaces + m.Name)
	spaces = spaces + "-----"

	if len(m.Children) == 0 {
		return
	}

	for _, c := range m.Children {
		Display(c, spaces)
	}
	fmt.Println()
}
