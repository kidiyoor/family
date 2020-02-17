package tejpal

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/kollu/tejpal/hithesh"
)

var (
  // Tejpal is a member of the family.
  Tejpal = types.Member{
    Name: "Tejpal",
    Children: []*types.Member{&hithesh.Hithesh, },
    Gender: types.Male,
  }
)
