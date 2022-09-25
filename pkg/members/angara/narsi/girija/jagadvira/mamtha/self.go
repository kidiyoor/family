package mamtha

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Mamtha is a member of the family.
  Mamtha = types.Member{
    Name: "Mamtha",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
