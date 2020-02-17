package deepali

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/gururaj/deepali/deeya"
)

var (
  // Deepali is a member of the family.
  Deepali = types.Member{
    Name: "Deepali",
    Children: []*types.Member{&deeya.Deeya, },
    Gender: types.Female,
  }
)
