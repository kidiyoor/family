package suneet

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Suneet is a member of the family.
  Suneet = types.Member{
    Name: "Suneet",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
