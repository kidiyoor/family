package vittal

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran/vittal/sumana"
)

var (
  // Vittal is a member of the family.
  Vittal = types.Member{
    Name: "Vittal",
    Children: []*types.Member{&sumana.Sumana, },
    Gender: types.Male,
  }
)
