package sanki

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/kalyani"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/lalitha"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/bhuvanendra"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/ganesh"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/sanjeevi"
)

var (
  // Sanki is a member of the family.
  Sanki = types.Member{
    Name: "Sanki",
    Children: []*types.Member{&bhuvanendra.Bhuvanendra, &ganesh.Ganesh, &kalyani.Kalyani, &lalitha.Lalitha, &sanjeevi.Sanjeevi, },
    Gender: types.Female,
  }
)
