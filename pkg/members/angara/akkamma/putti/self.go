package putti

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/putti/danapathi"
)

var (
  // Putti is a member of the family.
  Putti = types.Member{
    Name: "Putti",
    Children: []*types.Member{&danapathi.Danapathi, },
    Gender: types.Female,
  }
)
