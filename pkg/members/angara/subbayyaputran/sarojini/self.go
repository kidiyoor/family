package sarojini

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran/sarojini/g"
)

var (
  // Sarojini is a member of the family.
  Sarojini = types.Member{
    Name: "Sarojini",
    Children: []*types.Member{&g.G, },
    Gender: types.Female,
  }
)
