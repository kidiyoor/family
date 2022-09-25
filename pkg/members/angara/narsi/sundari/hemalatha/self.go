package hemalatha

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/hemalatha/divya"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/hemalatha/prerana"
)

var (
  // Hemalatha is a member of the family.
  Hemalatha = types.Member{
    Name: "Hemalatha",
    Children: []*types.Member{&divya.Divya, &prerana.Prerana, },
    Gender: types.Female,
  }
)
