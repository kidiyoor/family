package yogishdhar

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi/yogishdhar/poorna"
)

var (
  // Yogishdhar is a member of the family.
  Yogishdhar = types.Member{
    Name: "Yogishdhar",
    Children: []*types.Member{&poorna.Poorna, },
    Gender: types.Male,
  }
)
