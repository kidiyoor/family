package rohit

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Rohit is a member of the family.
  Rohit = types.Member{
    Name: "Rohit",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
