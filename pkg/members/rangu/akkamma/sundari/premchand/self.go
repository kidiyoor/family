package premchand

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/premchand/g"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/premchand/b"
)

var (
  // Premchand is a member of the family.
  Premchand = types.Member{
    Name: "Premchand",
    Children: []*types.Member{&b.B, &g.G, },
    Gender: types.Male,
  }
)
