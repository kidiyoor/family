package ashish

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Ashish is a member of the family.
  Ashish = types.Member{
    Name: "Ashish",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
