package narsi

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sharada"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/yashoda"
)

var (
  // Narsi is a member of the family.
  Narsi = types.Member{
    Name: "Narsi",
    Children: []*types.Member{&girija.Girija, &kamala.Kamala, &sharada.Sharada, &sundari.Sundari, &yashoda.Yashoda, },
    Gender: types.Female,
  }
)
