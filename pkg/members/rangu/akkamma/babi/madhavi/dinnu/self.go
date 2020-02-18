package dinnu

import (
	"kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi/dinnu/vignesh"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi/dinnu/vidhi"
)

var (
	// Dinnu is a member of the family.
	Dinnu = types.Member{
		Name:     "Dinesh",
    Children: []*types.Member{&vidhi.Vidhi, &vignesh.Vignesh, },
		Gender:   types.Male,
	}
)
