package jyothi

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/lalitha/jyothi/shishir"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/lalitha/jyothi/poorvi"
)

var (
  // Jyothi is a member of the family.
  Jyothi = types.Member{
    Name: "Jyothi",
    Children: []*types.Member{&poorvi.Poorvi, &shishir.Shishir, },
    Gender: types.Female,
  }
)
