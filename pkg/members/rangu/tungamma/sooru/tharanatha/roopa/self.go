package roopa

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Roopa is a member of the family.
  Roopa = types.Member{
    Name: "Roopa",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
