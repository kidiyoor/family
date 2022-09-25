package heera

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Heera is a member of the family.
  Heera = types.Member{
    Name: "Heera",
    Spouse: &types.Member{
        Name: "Uday Kumar",
        Gender: types.Male,
      },
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
