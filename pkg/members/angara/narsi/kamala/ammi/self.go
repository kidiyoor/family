package ammi

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/ammi/mahima"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/ammi/manvitha"
)

var (
  // Ammi is a member of the family.
  Ammi = types.Member{
    Name: "Ammi(Sumangala)",
    Children: []*types.Member{&mahima.Mahima, &manvitha.Manvitha, },
    Gender: types.Female,
  }
)
