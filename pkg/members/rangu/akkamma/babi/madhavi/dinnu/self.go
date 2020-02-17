package dinnu

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Dinnu is a member of the family.
  Dinnu = types.Member{
    Name: "Dinnu",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
