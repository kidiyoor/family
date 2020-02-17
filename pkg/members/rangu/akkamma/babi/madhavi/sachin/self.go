package sachin

import (
	"kidiyoor.io/family-tree/pkg/types"
)

var (
	// Sachin is a member of the family.
	Sachin = types.Member{
		Name:     "Sachin",
		Children: []*types.Member{},
		Gender:   types.Male,
	}
)
