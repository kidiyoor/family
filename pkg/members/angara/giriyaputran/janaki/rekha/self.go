package rekha

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Rekha is a member of the family.
  Rekha = types.Member{
    Name: "Rekha",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
