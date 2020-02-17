package sharada

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/sharada/sheela"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/sharada/leka"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki/sharada/gautham"
)

var (
  // Sharada is a member of the family.
  Sharada = types.Member{
    Name: "Sharada",
    Children: []*types.Member{&gautham.Gautham, &leka.Leka, &sheela.Sheela, },
    Gender: types.Female,
  }
)
