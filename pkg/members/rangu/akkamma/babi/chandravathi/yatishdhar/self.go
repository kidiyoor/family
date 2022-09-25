package yatishdhar

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi/yatishdhar/gaurav"
)

var (
  // Yatishdhar is a member of the family.
  Yatishdhar = types.Member{
    Name: "Yatishdhar",
    Spouse: &types.Member{
        Name: "Anjali",
        Gender: types.Female,
      },
    Children: []*types.Member{&gaurav.Gaurav, },
    Gender: types.Male,
  }
)
