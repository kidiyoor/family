package ujwala

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Ujwala is a member of the family.
  Ujwala = types.Member{
    Name: "Ujwala",
    Spouse: &types.Member{
        Name: "Ankit",
        Gender: types.Male,
      },
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
