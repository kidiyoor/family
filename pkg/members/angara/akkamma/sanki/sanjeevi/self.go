package sanjeevi

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/sanjeevi/g1"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/sanjeevi/g2"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/sanjeevi/kalavathi"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/sanjeevi/b1"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/sanjeevi/b2"
)

var (
  // Sanjeevi is a member of the family.
  Sanjeevi = types.Member{
    Name: "Sanjeevi",
    Children: []*types.Member{&b1.B1, &b2.B2, &g1.G1, &g2.G2, &kalavathi.Kalavathi, },
    Gender: types.Female,
  }
)
