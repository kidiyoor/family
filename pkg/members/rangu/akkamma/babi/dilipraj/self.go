package dilipraj

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/dilipraj/ujwala"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/dilipraj/mridulla"
)

var (
  // Dilipraj is a member of the family.
  Dilipraj = types.Member{
    Name: "Dilipraj",
    Spouse: &types.Member{
        Name: "Jayanthi",
        Gender: types.Female,
      },
    Children: []*types.Member{&mridulla.Mridulla, &ujwala.Ujwala, },
    Gender: types.Male,
  }
)
