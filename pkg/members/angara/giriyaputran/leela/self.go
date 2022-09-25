package leela

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/leela/b1"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/leela/b2"
)

var (
  // Leela is a member of the family.
  Leela = types.Member{
    Name: "Leela",
    Children: []*types.Member{&b1.B1, &b2.B2, },
    Gender: types.Female,
  }
)
