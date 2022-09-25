package kamala

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran/kamala/sathya"
)

var (
  // Kamala is a member of the family.
  Kamala = types.Member{
    Name: "Kamala",
    Children: []*types.Member{&sathya.Sathya, },
    Gender: types.Female,
  }
)
