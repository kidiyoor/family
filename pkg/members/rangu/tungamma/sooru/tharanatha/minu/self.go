package minu

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Minu is a member of the family.
  Minu = types.Member{
    Name: "Minu",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
