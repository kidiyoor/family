package madhavi

import (
	"kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi/bipin"
	"kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi/dinnu"
	"kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi/sachin"
	"kidiyoor.io/family-tree/pkg/types"
)

var (
	// Madhavi is a member of the family.
	Madhavi = types.Member{
		Name:     "Madhavi",
		Children: []*types.Member{&bipin.Bipin, &dinnu.Dinnu, &sachin.Sachin},
		Gender:   types.Female,
	}
)
