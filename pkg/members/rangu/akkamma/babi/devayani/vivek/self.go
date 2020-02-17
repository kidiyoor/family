package vivek

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Vivek is a member of the family.
  Vivek = types.Member{
    Name: "Vivek",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
