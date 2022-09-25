package vineet

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Vineet is a member of the family.
  Vineet = types.Member{
    Name: "Vineet",
    Spouse: &types.Member{
        Name: "Nisha",
        Gender: types.Female,
      },
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
