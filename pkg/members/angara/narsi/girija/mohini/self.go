package mohini

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/mohini/manish"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/mohini/nithin"
)

var (
  // Mohini is a member of the family.
  Mohini = types.Member{
    Name: "Mohini",
    Children: []*types.Member{&manish.Manish, &nithin.Nithin, },
    Gender: types.Female,
  }
)
