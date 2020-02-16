package shesha_sahukar

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/shesha_sahukar/kamala"
  "kidiyoor.io/family-tree/pkg/members/rangu/shesha_sahukar/muthakka"
  "kidiyoor.io/family-tree/pkg/members/rangu/shesha_sahukar/nagaveni"
  "kidiyoor.io/family-tree/pkg/members/rangu/shesha_sahukar/neelamma"
  "kidiyoor.io/family-tree/pkg/members/rangu/shesha_sahukar/sharadha"
)

var (
	Shesha_sahukar = types.Member{
		Name: "Shesha sahukar",
    Children: []*types.Member{
      &neelamma.Neelamma, &muthakka.Muthakka, &kamala.Kamala, &nagaveni.Nagaveni, &sharadha.Sharadha,
    },
	}
)
