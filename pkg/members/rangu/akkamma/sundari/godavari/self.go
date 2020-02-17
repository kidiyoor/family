package godavari

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/godavari/sandhya"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/godavari/chandini"
)

var (
  // Godavari is a member of the family.
  Godavari = types.Member{
    Name: "Godavari",
    Children: []*types.Member{&chandini.Chandini, &sandhya.Sandhya, },
    Gender: types.Female,
  }
)
