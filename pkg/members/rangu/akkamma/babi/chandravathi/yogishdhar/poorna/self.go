package poorna

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Poorna is a member of the family.
  Poorna = types.Member{
    Name: "Poorna",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
