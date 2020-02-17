package rahul

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Rahul is a member of the family.
  Rahul = types.Member{
    Name: "Rahul",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
