package manish

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Manish is a member of the family.
  Manish = types.Member{
    Name: "Manish",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
