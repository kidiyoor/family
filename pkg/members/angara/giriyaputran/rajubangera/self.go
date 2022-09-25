package rajubangera

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/rajubangera/b"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/rajubangera/g"
)

var (
  // Rajubangera is a member of the family.
  Rajubangera = types.Member{
    Name: "RajuBangera",
    Children: []*types.Member{&b.B, &g.G, },
    Gender: types.Male,
  }
)
