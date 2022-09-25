package angara

import (
	"kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran"
  "kidiyoor.io/family-tree/pkg/members/angara/krishnappakidiyoor"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi"
  "kidiyoor.io/family-tree/pkg/members/angara/taniyaajjer"
)

var (
	// Angara is a family member.
	Angara = types.Member{
		Name:     "Angara",
		Gender:   types.Male,
    Children: []*types.Member{&akkamma.Akkamma, &giriyaputran.Giriyaputran, &krishnappakidiyoor.Krishnappakidiyoor, &narsi.Narsi, &subbayyaputran.Subbayyaputran, &taniyaajjer.Taniyaajjer, },
	}
)
