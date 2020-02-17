package girish

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Girish is a member of the family.
  Girish = types.Member{
    Name: "Girish",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
