package koragappa

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/koragappa/gangadhar"
  "kidiyoor.io/family-tree/pkg/members/rangu/koragappa/chinnapa"
  "kidiyoor.io/family-tree/pkg/members/rangu/koragappa/indira"
  "kidiyoor.io/family-tree/pkg/members/rangu/koragappa/sadashiva"
  "kidiyoor.io/family-tree/pkg/members/rangu/koragappa/vimala"
  "kidiyoor.io/family-tree/pkg/members/rangu/koragappa/kamalaksha"
  "kidiyoor.io/family-tree/pkg/members/rangu/koragappa/yashoda"
)

var (
  // Koragappa is a member of the family.
  Koragappa = types.Member{
    Name: "Koragappa",
    Children: []*types.Member{&chinnapa.Chinnapa, &gangadhar.Gangadhar, &indira.Indira, &kamalaksha.Kamalaksha, &sadashiva.Sadashiva, &vimala.Vimala, &yashoda.Yashoda, },
    Gender: types.Male,
  }
)
