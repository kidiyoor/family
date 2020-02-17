package geetha

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/geetha/prashant"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/geetha/veena"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/geetha/chitra"
)

var (
  // Geetha is a member of the family.
  Geetha = types.Member{
    Name: "Geetha",
    Children: []*types.Member{&chitra.Chitra, &prashant.Prashant, &veena.Veena, },
    Gender: types.Female,
  }
)
