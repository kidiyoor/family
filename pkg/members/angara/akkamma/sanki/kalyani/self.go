package kalyani

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/kalyani/sukumar"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/kalyani/pushpa"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/kalyani/nalini"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/kalyani/surekha"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/kalyani/raja"
)

var (
  // Kalyani is a member of the family.
  Kalyani = types.Member{
    Name: "Kalyani",
    Children: []*types.Member{&nalini.Nalini, &pushpa.Pushpa, &raja.Raja, &sukumar.Sukumar, &surekha.Surekha, },
    Gender: types.Female,
  }
)
