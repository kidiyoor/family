package sharada

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/sharada/ashish"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/sharada/devanand"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/sharada/mohini"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/sharada/shalinia"
)

var (
  // Sharada is a member of the family.
  Sharada = types.Member{
    Name: "Sharada",
    Children: []*types.Member{&ashish.Ashish, &devanand.Devanand, &mohini.Mohini, &shalinia.Shalinia, },
    Gender: types.Female,
  }
)
