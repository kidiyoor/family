package chandrahasa

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/chandrahasa/chethan"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/chandrahasa/roshan"
)

var (
  // Chandrahasa is a member of the family.
  Chandrahasa = types.Member{
    Name: "Chandrahasa",
    Children: []*types.Member{&chethan.Chethan, &roshan.Roshan, },
    Gender: types.Male,
  }
)
