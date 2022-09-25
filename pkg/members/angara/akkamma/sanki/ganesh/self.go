package ganesh

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki/ganesh/g"
)

var (
  // Ganesh is a member of the family.
  Ganesh = types.Member{
    Name: "Ganesh",
    Children: []*types.Member{&g.G, },
    Gender: types.Male,
  }
)
