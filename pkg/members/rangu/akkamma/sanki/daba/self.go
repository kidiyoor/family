package daba

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/daba/mamatha"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/daba/santhosh"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/daba/b"
)

var (
  // Daba is a member of the family.
  Daba = types.Member{
    Name: "Daba",
    Children: []*types.Member{&b.B, &mamatha.Mamatha, &santhosh.Santhosh, },
    Gender: types.Male,
  }
)
