package devayani

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/devayani/vivek"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/devayani/mamta"
)

var (
  // Devayani is a member of the family.
  Devayani = types.Member{
    Name: "Devayani",
    Children: []*types.Member{&mamta.Mamta, &vivek.Vivek, },
    Gender: types.Female,
  }
)
