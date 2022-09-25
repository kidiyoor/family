package vinod

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Vinod is a member of the family.
  Vinod = types.Member{
    Name: "Vinod",
    Children: []*types.Member{},
    Gender: types.Male,
  }
)
