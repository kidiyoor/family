package hemagiri

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Hemagiri is a member of the family.
  Hemagiri = types.Member{
    Name: "hemagiri",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
