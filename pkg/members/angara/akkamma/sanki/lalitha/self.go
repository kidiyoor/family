package lalitha

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/lalitha/b"
)

var (
  // Lalitha is a member of the family.
  Lalitha = types.Member{
    Name: "Lalitha",
    Children: []*types.Member{&b.B, },
    Gender: types.Female,
  }
)
