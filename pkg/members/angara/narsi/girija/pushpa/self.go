package pushpa

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Pushpa is a member of the family.
  Pushpa = types.Member{
    Name: "Pushpa",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
