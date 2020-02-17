package veena

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Veena is a member of the family.
  Veena = types.Member{
    Name: "Veena",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
