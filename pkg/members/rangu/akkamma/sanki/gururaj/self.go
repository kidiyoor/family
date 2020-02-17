package gururaj

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/gururaj/harshali"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/gururaj/deepali"
)

var (
  // Gururaj is a member of the family.
  Gururaj = types.Member{
    Name: "Gururaj",
    Children: []*types.Member{&deepali.Deepali, &harshali.Harshali, },
    Gender: types.Male,
  }
)
