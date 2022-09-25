package bipin

import (
	"kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi/bipin/premraj"
)

var (
	// Bipin is a member of the family.
	Bipin = types.Member{
		Name:     "Bipinchandra",
    Spouse: &types.Member{
        Name: "Asha",
        Gender: types.Female,
      },
    Children: []*types.Member{&premraj.Premraj, },
		Gender:   types.Male,
	}
)
