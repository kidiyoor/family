package nithin

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/harini/nithin/sanisha"
)

var (
  // Nithin is a member of the family.
  Nithin = types.Member{
    Name: "Nithin",
    Children: []*types.Member{&sanisha.Sanisha, },
    Gender: types.Male,
  }
)
