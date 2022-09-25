package raju

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Raju is a member of the family.
  Raju = types.Member{
    Name: "Raju",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
