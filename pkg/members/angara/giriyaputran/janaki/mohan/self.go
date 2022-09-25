package mohan

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Mohan is a member of the family.
  Mohan = types.Member{
    Name: "Mohan",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
