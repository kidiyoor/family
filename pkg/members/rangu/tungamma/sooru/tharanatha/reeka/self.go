package reeka

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Reeka is a member of the family.
  Reeka = types.Member{
    Name: "Reeka",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
