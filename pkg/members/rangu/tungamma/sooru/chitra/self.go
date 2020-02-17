package chitra

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/chitra/shuba"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/chitra/prashant"
)

var (
  // Chitra is a member of the family.
  Chitra = types.Member{
    Name: "Chitra",
    Children: []*types.Member{&prashant.Prashant, &shuba.Shuba, },
    Gender: types.Female,
  }
)
