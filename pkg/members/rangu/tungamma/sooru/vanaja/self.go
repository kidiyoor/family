package vanaja

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/vanaja/rohit"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/vanaja/devaki"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/vanaja/giri"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/vanaja/sarala"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/vanaja/indira"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/vanaja/shesha"
)

var (
  // Vanaja is a member of the family.
  Vanaja = types.Member{
    Name: "Vanaja",
    Children: []*types.Member{&devaki.Devaki, &giri.Giri, &indira.Indira, &rohit.Rohit, &sarala.Sarala, &shesha.Shesha, },
    Gender: types.Female,
  }
)
