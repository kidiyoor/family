package lokanath

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/lokanath/prajna"
)

var (
  // Lokanath is a member of the family.
  Lokanath = types.Member{
    Name: "Lokanath",
    Spouse: &types.Member{
        Name: "Pramila",
        Gender: types.Female,
      },
    Children: []*types.Member{&prajna.Prajna, },
    Gender: types.Male,
  }
)
