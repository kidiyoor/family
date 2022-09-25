package pramila

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/pramila/prajna"
)

var (
  // Pramila is a member of the family.
  Pramila = types.Member{
    Name: "Pramila",
    Children: []*types.Member{&prajna.Prajna, },
    Gender: types.Female,
  }
)
