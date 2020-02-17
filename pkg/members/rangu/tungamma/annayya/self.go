package annayya

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/annayya/bharathi"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/annayya/pushpa"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/annayya/raviraja"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/annayya/sadananda"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/annayya/g"
)

var (
  // Annayya is a member of the family.
  Annayya = types.Member{
    Name: "Annayya",
    Children: []*types.Member{&bharathi.Bharathi, &g.G, &pushpa.Pushpa, &raviraja.Raviraja, &sadananda.Sadananda, },
    Gender: types.Male,
  }
)
