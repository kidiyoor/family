package ramappa

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa/hemavathi"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa/kasturi"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa/goverdan"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa/jhanavi"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa/preema"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa/somanatha"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa/jayaprakash"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa/yadavi"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa/bojaraj"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa/girraji"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa/anusuya"
)

var (
  // Ramappa is a member of the family.
  Ramappa = types.Member{
    Name: "Ramappa",
    Children: []*types.Member{&anusuya.Anusuya, &bojaraj.Bojaraj, &girraji.Girraji, &goverdan.Goverdan, &hemavathi.Hemavathi, &jayaprakash.Jayaprakash, &jhanavi.Jhanavi, &kasturi.Kasturi, &preema.Preema, &somanatha.Somanatha, &yadavi.Yadavi, },
    Gender: types.Male,
  }
)
