package kiara

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Kiara is a member of the family.
  Kiara = types.Member{
    Name: "Kiara",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
