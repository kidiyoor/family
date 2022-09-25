package rama

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Rama is a member of the family.
  Rama = types.Member{
    Name: "Rama",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
