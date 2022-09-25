package mamtha

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/mamtha/shravya"
)

var (
  // Mamtha is a member of the family.
  Mamtha = types.Member{
    Name: "Mamtha",
    Children: []*types.Member{&shravya.Shravya, },
    Gender: types.Female,
  }
)
