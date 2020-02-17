package kollu

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/kollu/gopi"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/kollu/devu"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/kollu/tejpal"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/kollu/mohini"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/kollu/yeshoda"
)

var (
  // Kollu is a member of the family.
  Kollu = types.Member{
    Name: "Kollu",
    Children: []*types.Member{&devu.Devu, &gopi.Gopi, &mohini.Mohini, &tejpal.Tejpal, &yeshoda.Yeshoda, },
    Gender: types.Female,
  }
)
