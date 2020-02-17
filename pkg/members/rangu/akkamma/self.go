package akkamma

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/kollu"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sanki"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/mahabala"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari"
)

var (
  // Akkamma is a member of the family.
  Akkamma = types.Member{
    Name: "Akkamma",
    Children: []*types.Member{&babi.Babi, &kollu.Kollu, &mahabala.Mahabala, &sanki.Sanki, &sundari.Sundari, },
    Gender: types.Female,
  }
)
