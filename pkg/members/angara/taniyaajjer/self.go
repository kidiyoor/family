package taniyaajjer

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/taniyaajjer/madava"
  "kidiyoor.io/family-tree/pkg/members/angara/taniyaajjer/raampe"
)

var (
  // Taniyaajjer is a member of the family.
  Taniyaajjer = types.Member{
    Name: "TaniyaAjjer",
    Children: []*types.Member{&madava.Madava, &raampe.Raampe, },
    Gender: types.Male,
  }
)
