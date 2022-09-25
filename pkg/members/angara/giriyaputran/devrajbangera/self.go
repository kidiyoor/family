package devrajbangera

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/devrajbangera/b1"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/devrajbangera/b2"
)

var (
  // Devrajbangera is a member of the family.
  Devrajbangera = types.Member{
    Name: "DevrajBangera",
    Children: []*types.Member{&b1.B1, &b2.B2, },
    Gender: types.Male,
  }
)
