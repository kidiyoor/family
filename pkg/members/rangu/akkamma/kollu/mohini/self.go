package mohini

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/kollu/mohini/girish"
)

var (
  // Mohini is a member of the family.
  Mohini = types.Member{
    Name: "Mohini",
    Children: []*types.Member{&girish.Girish, },
    Gender: types.Female,
  }
)
