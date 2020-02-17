package yeshoda

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Yeshoda is a member of the family.
  Yeshoda = types.Member{
    Name: "Yeshoda",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
