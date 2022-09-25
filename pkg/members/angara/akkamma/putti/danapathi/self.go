package danapathi

import (
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran/kamala/sathya"
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Danapathi is a member of the family.
  Danapathi = types.Member{
    Name: "Danapathi",
    Spouse: &sathya.Sathya,
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
