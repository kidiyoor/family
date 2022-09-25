package kamal

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Kamal is a member of the family.
  Kamal = types.Member{
    Name: "Kamal",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
