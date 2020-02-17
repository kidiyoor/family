package devu

import (
	"kidiyoor.io/family-tree/pkg/members/rangu/akkamma/kollu/devu/suraj"
	"kidiyoor.io/family-tree/pkg/types"
)

var (
	// Devu is a member of the family.
	Devu = types.Member{
		Name:     "Devu",
		Children: []*types.Member{&suraj.Suraj},
		Gender:   types.Male,
	}
)
