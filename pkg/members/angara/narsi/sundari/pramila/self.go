package pramila

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/lokanath/prajna"
)

var (
  // Pramila is a member of the family.
  Pramila = types.Member{
    Name: "Pramila",
    Children: []*types.Member{&prajna.Prajna, },
    Gender: types.Female,
  }
)
