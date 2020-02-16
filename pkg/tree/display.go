package tree

import (
	"kidiyoor.io/family-tree/pkg/members"
)

// Display returns the family tree in string format.
func Display() string {
	str := ""
	for _, m := range members.Root.Children {
		str = str + " | " + m.Name
	}

	return str
}
