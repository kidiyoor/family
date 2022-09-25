package bhuvanendra

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/bhuvanendra/b1"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/bhuvanendra/b2"
)

var (
  // Bhuvanendra is a member of the family.
  Bhuvanendra = types.Member{
    Name: "Bhuvanendra",
    Children: []*types.Member{&b1.B1, &b2.B2, },
    Gender: types.Male,
  }
)
