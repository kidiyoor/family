package preveena

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/preveena/b1"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/preveena/b2"
)

var (
  // Preveena is a member of the family.
  Preveena = types.Member{
    Name: "Preveena",
    Children: []*types.Member{&b1.B1, &b2.B2, },
    Gender: types.Female,
  }
)
