package tree

import (
	"fmt"

	"kidiyoor.io/family-tree/pkg/types"
)

// Display the family tree.
func Display(m *types.Member, spaces string) {
        couple := m.Name
	if m.Spouse != nil {
             couple = couple + " <3 " + m.Spouse.Name
        }
        fmt.Println(spaces + couple)
        spaces = spaces + "--------"

	if len(m.Children) == 0 {
		return
	}

	for _, c := range m.Children {
		Display(c, spaces)
	}
	fmt.Println()
}
