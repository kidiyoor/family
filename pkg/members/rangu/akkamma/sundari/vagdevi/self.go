package vagdevi

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/vagdevi/deepak"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/vagdevi/divya"
)

var (
  // Vagdevi is a member of the family.
  Vagdevi = types.Member{
    Name: "Vagdevi",
    Children: []*types.Member{&deepak.Deepak, &divya.Divya, },
    Gender: types.Female,
  }
)
