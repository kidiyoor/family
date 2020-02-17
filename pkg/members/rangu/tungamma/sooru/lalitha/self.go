package lalitha

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/lalitha/shoba"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/lalitha/jyothi"
)

var (
  // Lalitha is a member of the family.
  Lalitha = types.Member{
    Name: "Lalitha",
    Children: []*types.Member{&jyothi.Jyothi, &shoba.Shoba, },
    Gender: types.Female,
  }
)
