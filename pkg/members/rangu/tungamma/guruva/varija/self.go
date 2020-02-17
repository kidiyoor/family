package varija

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Varija is a member of the family.
  Varija = types.Member{
    Name: "Varija",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
