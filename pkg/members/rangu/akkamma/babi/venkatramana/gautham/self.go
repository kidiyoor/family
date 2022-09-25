package gautham

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Gautham is a member of the family.
  Gautham = types.Member{
    Name: "Gautham",
    Spouse: &types.Member{
        Name: "Swarna Sathya",
        Gender: types.Female,
      },
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
