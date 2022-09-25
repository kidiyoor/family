package chandravathi

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi/girishdhar"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi/yatishdhar"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi/yogishdhar"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi/heera"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi/dhavalesh"
)

var (
  // Chandravathi is a member of the family.
  Chandravathi = types.Member{
    Name: "Chandravathi",
    Spouse: &types.Member{
        Name: "Gangadhar",
        Gender: types.Male,
      },
    Children: []*types.Member{&dhavalesh.Dhavalesh, &girishdhar.Girishdhar, &heera.Heera, &yatishdhar.Yatishdhar, &yogishdhar.Yogishdhar, },
    Gender: types.Female,
  }
)
