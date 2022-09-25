package vivek

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/devayani/vivek/kiara"
)

var (
  // Vivek is a member of the family.
  Vivek = types.Member{
    Name: "Vivek",
    Spouse: &types.Member{
        Name: "Harsha",
        Gender: types.Female,
      },
    Children: []*types.Member{&kiara.Kiara, },
    Gender: types.Male,
  }
)
