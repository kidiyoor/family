package dhavalesh

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi/dhavalesh/sanath"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi/dhavalesh/siddhan"
)

var (
  // Dhavalesh is a member of the family.
  Dhavalesh = types.Member{
    Name: "Dhavalesh",
    Spouse: &types.Member{
        Name: "Mallika",
        Gender: types.Female,
      },
    Children: []*types.Member{&sanath.Sanath, &siddhan.Siddhan, },
    Gender: types.Male,
  }
)
