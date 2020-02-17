package rangu

import (
	"kidiyoor.io/family-tree/pkg/members/rangu/akkamma"
	"kidiyoor.io/family-tree/pkg/members/rangu/chaudappa"
	"kidiyoor.io/family-tree/pkg/members/rangu/koragappa"
	"kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma"
	"kidiyoor.io/family-tree/pkg/types"
)

var (
	// Rangu is a family member.
	Rangu = types.Member{
		Name:     "Rangu",
		Gender:   types.Male,
		Children: []*types.Member{&akkamma.Akkamma, &chaudappa.Chaudappa, &koragappa.Koragappa, &sheshasahukar.Sheshasahukar, &tungamma.Tungamma},
	}
)
