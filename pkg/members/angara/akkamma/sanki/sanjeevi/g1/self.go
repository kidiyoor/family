package g1

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // G1 is a member of the family.
  G1 = types.Member{
    Name: "g1",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
