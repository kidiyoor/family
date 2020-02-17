package yashoda

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Yashoda is a member of the family.
  Yashoda = types.Member{
    Name: "Yashoda",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
