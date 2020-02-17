package divya

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Divya is a member of the family.
  Divya = types.Member{
    Name: "Divya",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
