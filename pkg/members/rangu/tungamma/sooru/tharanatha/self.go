package tharanatha

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/tharanatha/roopa"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/tharanatha/reeka"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/tharanatha/b"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/tharanatha/minu"
)

var (
  // Tharanatha is a member of the family.
  Tharanatha = types.Member{
    Name: "Tharanatha",
    Children: []*types.Member{&b.B, &minu.Minu, &reeka.Reeka, &roopa.Roopa, },
    Gender: types.Male,
  }
)
