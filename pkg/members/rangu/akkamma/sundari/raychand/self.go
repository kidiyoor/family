package raychand

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/raychand/pratap"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/raychand/prajwal"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/raychand/pradeep"
)

var (
  // Raychand is a member of the family.
  Raychand = types.Member{
    Name: "Raychand",
    Children: []*types.Member{&pradeep.Pradeep, &prajwal.Prajwal, &pratap.Pratap, },
    Gender: types.Male,
  }
)
