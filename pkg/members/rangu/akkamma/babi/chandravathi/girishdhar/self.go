package girishdhar

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi/girishdhar/greeshma"
)

var (
  // Girishdhar is a member of the family.
  Girishdhar = types.Member{
    Name: "Girishdhar",
    Children: []*types.Member{&greeshma.Greeshma, },
    Gender: types.Male,
  }
)
