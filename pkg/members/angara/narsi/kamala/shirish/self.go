package shirish

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/shirish/yanisha"
)

var (
  // Shirish is a member of the family.
  Shirish = types.Member{
    Name: "Shirish",
    Children: []*types.Member{&yanisha.Yanisha, },
    Gender: types.Male,
  }
)
