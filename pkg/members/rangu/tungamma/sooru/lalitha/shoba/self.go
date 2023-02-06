package shoba

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/lalitha/shoba/sharvil"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/lalitha/shoba/sharanya"
)

var (
  // Shoba is a member of the family.
  Shoba = types.Member{
    Name: "Shoba & Dinkar",
    Children: []*types.Member{&sharanya.Sharanya, &sharvil.Sharvil, },
    Gender: types.Female,
  }
)
