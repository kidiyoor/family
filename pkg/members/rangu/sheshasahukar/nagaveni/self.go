package nagaveni

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/nagaveni/tharachuda"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/nagaveni/damodhara"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/nagaveni/keshava"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/nagaveni/kunnudimi"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/nagaveni/lalaji"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/nagaveni/g"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/nagaveni/dinakar"
)

var (
  // Nagaveni is a member of the family.
  Nagaveni = types.Member{
    Name: "Nagaveni",
    Children: []*types.Member{&damodhara.Damodhara, &dinakar.Dinakar, &g.G, &keshava.Keshava, &kunnudimi.Kunnudimi, &lalaji.Lalaji, &tharachuda.Tharachuda, },
    Gender: types.Female,
  }
)
