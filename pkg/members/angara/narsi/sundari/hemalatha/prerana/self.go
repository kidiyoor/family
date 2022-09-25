package prerana

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/hemalatha/prerana/divesh"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/hemalatha/prerana/g"
)

var (
  // Prerana is a member of the family.
  Prerana = types.Member{
    Name: "Prerana",
    Children: []*types.Member{&divesh.Divesh, &g.G, },
    Gender: types.Female,
  }
)
