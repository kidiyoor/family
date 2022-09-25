package soumya

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/soumya/spoorthi"
)

var (
  // Soumya is a member of the family.
  Soumya = types.Member{
    Name: "Soumya",
    Children: []*types.Member{&spoorthi.Spoorthi, },
    Gender: types.Female,
  }
)
