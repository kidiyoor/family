package deepak

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Deepak is a member of the family.
  Deepak = types.Member{
    Name: "Deepak",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
