package daskshayani

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/daskshayani/rahul"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/daskshayani/g"
)

var (
  // Daskshayani is a member of the family.
  Daskshayani = types.Member{
    Name: "Daskshayani",
    Children: []*types.Member{&g.G, &rahul.Rahul, },
    Gender: types.Female,
  }
)
