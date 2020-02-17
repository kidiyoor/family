package sharada

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/sharada/vidya"
)

var (
  // Sharada is a member of the family.
  Sharada = types.Member{
    Name: "Sharada",
    Children: []*types.Member{&vidya.Vidya, },
    Gender: types.Female,
  }
)
