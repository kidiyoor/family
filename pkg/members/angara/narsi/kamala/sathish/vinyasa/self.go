package vinyasa

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Vinyasa is a member of the family.
  Vinyasa = types.Member{
    Name: "Vinyasa",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
