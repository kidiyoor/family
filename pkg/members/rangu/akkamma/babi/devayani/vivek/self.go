package vivek

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/devayani/vivek/kiara"
)

var (
  // Vivek is a member of the family.
  Vivek = types.Member{
    Name: "Vivek",
    Children: []*types.Member{&kiara.Kiara, },
    Gender: types.Male,
  }
)
