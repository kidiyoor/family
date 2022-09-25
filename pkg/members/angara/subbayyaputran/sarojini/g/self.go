package g

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // G is a member of the family.
  G = types.Member{
    Name: "g",
    Spouse: &types.Member{
        Name: "Srinivas",
        Gender: types.Male,
      },
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
