package b

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // B is a member of the family.
  B = types.Member{
    Name: "B",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
