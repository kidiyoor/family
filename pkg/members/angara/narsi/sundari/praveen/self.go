package praveen

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Praveen is a member of the family.
  Praveen = types.Member{
    Name: "Praveen",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
