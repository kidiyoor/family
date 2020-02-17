package sanki

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/sharada"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/daba"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/gururaj"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/madhava"
)

var (
  // Sanki is a member of the family.
  Sanki = types.Member{
    Name: "Sanki",
    Children: []*types.Member{&daba.Daba, &gururaj.Gururaj, &madhava.Madhava, &sharada.Sharada, },
    Gender: types.Female,
  }
)
