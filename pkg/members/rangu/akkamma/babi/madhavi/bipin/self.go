package bipin

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Bipin is a member of the family.
  Bipin = types.Member{
    Name: "Bipin",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
