package jaya

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Jaya is a member of the family.
  Jaya = types.Member{
    Name: "Jaya",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
