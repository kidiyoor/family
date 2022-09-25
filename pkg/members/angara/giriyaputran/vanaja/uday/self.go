package uday

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Uday is a member of the family.
  Uday = types.Member{
    Name: "Uday",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
