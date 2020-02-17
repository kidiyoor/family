package kamala

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/kamala/dhanalaxmipurandara"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/kamala/heera"
)

var (
  // Kamala is a member of the family.
  Kamala = types.Member{
    Name: "Kamala",
    Children: []*types.Member{&dhanalaxmipurandara.Dhanalaxmipurandara, &heera.Heera, },
    Gender: types.Female,
  }
)
