package shesha

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Shesha is a member of the family.
  Shesha = types.Member{
    Name: "Shesha",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
