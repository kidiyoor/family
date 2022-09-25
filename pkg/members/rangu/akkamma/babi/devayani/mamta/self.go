package mamta

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/devayani/mamta/nihit"
)

var (
  // Mamta is a member of the family.
  Mamta = types.Member{
    Name: "Mamta",
    Spouse: &types.Member{
        Name: "Jayaram",
        Gender: types.Male,
      },
    Children: []*types.Member{&nihit.Nihit, },
    Gender: types.Female,
  }
)
