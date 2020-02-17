package chandraprabha

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/chandraprabha/b1"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/chandraprabha/b2"
)

var (
  // Chandraprabha is a member of the family.
  Chandraprabha = types.Member{
    Name: "Chandraprabha",
    Children: []*types.Member{&b1.B1, &b2.B2, },
    Gender: types.Female,
  }
)
