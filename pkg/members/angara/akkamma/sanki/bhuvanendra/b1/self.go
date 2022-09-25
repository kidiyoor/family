package b1

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // B1 is a member of the family.
  B1 = types.Member{
    Name: "b1",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
