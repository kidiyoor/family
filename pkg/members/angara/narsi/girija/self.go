package girija

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/mohini"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/jagadvira"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/kasturi"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/pushpa"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/chandrahasa"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/soumya"
)

var (
  // Girija is a member of the family.
  Girija = types.Member{
    Name: "Girija",
    Children: []*types.Member{&chandrahasa.Chandrahasa, &jagadvira.Jagadvira, &kasturi.Kasturi, &mohini.Mohini, &pushpa.Pushpa, &soumya.Soumya, },
    Gender: types.Female,
  }
)
