package mamta

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Mamta is a member of the family.
  Mamta = types.Member{
    Name: "Mamta",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
