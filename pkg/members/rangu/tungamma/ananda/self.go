package ananda

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ananda/divakar"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ananda/dinakar"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ananda/sevanthi"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ananda/geetha"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ananda/gayathri"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ananda/g"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ananda/swarna"
)

var (
  // Ananda is a member of the family.
  Ananda = types.Member{
    Name: "Ananda",
    Children: []*types.Member{&dinakar.Dinakar, &divakar.Divakar, &g.G, &gayathri.Gayathri, &geetha.Geetha, &sevanthi.Sevanthi, &swarna.Swarna, },
    Gender: types.Male,
  }
)
