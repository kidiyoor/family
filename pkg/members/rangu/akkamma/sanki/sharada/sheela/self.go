package sheela

import (
	"kidiyoor.io/family-tree/pkg/types"
)

var (
	// Sheela is a member of the family.
	Sheela = types.Member{
		Name:     "Sheela",
		Children: []*types.Member{},
		Gender:   types.Female,
	}
)
