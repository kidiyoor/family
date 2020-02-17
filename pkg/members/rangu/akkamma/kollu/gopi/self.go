package gopi

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/kollu/gopi/venugopal"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/kollu/gopi/sabitha"
)

var (
  // Gopi is a member of the family.
  Gopi = types.Member{
    Name: "Gopi",
    Children: []*types.Member{&sabitha.Sabitha, &venugopal.Venugopal, },
    Gender: types.Female,
  }
)
