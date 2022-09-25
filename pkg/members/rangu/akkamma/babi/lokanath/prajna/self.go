package prajna

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/lokanath/prajna/myra"
)

var (
  // Prajna is a member of the family.
  Prajna = types.Member{
    Name: "Prajna",
    Spouse: &types.Member{
        Name: "Peter",
        Gender: types.Male,
      },
    Children: []*types.Member{&myra.Myra, },
    Gender: types.Female,
  }
)
