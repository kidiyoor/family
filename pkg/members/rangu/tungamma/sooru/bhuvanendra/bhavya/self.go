package bhavya

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Bhavya is a member of the family.
  Bhavya = types.Member{
    Name: "Bhavya",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
