package roshan

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Roshan is a member of the family.
  Roshan = types.Member{
    Name: "Roshan",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
