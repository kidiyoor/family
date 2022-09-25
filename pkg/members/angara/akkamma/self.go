package akkamma

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/putti"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/sanki"
  "kidiyoor.io/family-tree/pkg/members/angara/akkamma/venkatesh"
)

var (
  // Akkamma is a member of the family.
  Akkamma = types.Member{
    Name: "Akkamma",
    Children: []*types.Member{&putti.Putti, &sanki.Sanki, &venkatesh.Venkatesh, },
    Gender: types.Female,
  }
)
