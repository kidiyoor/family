package kalavathi

import (
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/kalavathi/geetha"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/kalavathi/jaya"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/kalavathi/santosh"
	"kidiyoor.io/family-tree/pkg/types"
)

var (
	// Kalavathi is a member of the family.
	Kalavathi = types.Member{
		Name:     "Kalavathi",
		Children: []*types.Member{&geetha.Geetha, &jaya.Jaya, &santosh.Santosh},
		Gender:   types.Female,
	}
)
