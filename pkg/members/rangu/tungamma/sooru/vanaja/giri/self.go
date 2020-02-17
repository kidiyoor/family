package giri

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Giri is a member of the family.
  Giri = types.Member{
    Name: "Giri",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
