package sachin

import (
	"kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi/sachin/rahul"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi/sachin/dhrithi"
)

var (
	// Sachin is a member of the family.
	Sachin = types.Member{
		Name:     "Sachinchandra",
    Spouse: &types.Member{
        Name: "Neeta",
        Gender: types.Female,
      },
    Children: []*types.Member{&dhrithi.Dhrithi, &rahul.Rahul, },
		Gender:   types.Male,
	}
)
