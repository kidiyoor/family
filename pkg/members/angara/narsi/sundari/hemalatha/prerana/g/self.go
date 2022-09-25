package g

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // G is a member of the family.
  G = types.Member{
    Name: "g",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
