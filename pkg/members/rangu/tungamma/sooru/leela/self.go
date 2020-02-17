package leela

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/leela/sujith"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/leela/vineet"
)

var (
  // Leela is a member of the family.
  Leela = types.Member{
    Name: "Leela",
    Children: []*types.Member{&sujith.Sujith, &vineet.Vineet, },
    Gender: types.Female,
  }
)
