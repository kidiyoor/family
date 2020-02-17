package madhava

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/madhava/vani"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/madhava/veena"
)

var (
  // Madhava is a member of the family.
  Madhava = types.Member{
    Name: "Madhava",
    Children: []*types.Member{&vani.Vani, &veena.Veena, },
    Gender: types.Male,
  }
)
