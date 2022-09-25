package yogishdhar

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi/yogishdhar/poorna"
)

var (
  // Yogishdhar is a member of the family.
  Yogishdhar = types.Member{
    Name: "Yogishdhar",
    Spouse: &types.Member{
        Name: "Mallika",
        Gender: types.Female,
      },
    Children: []*types.Member{&poorna.Poorna, },
    Gender: types.Male,
  }
)
