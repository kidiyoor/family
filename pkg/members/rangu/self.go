package rangu

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma"
  "kidiyoor.io/family-tree/pkg/members/rangu/choudappa"
  "kidiyoor.io/family-tree/pkg/members/rangu/koragappa"
  "kidiyoor.io/family-tree/pkg/members/rangu/shesha_sahukar"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma"
)

var (
	Rangu = types.Member{
		Name:   "Rangu",
		Gender: types.Male,
		Children: []*types.Member{
			&shesha_sahukar.Shesha_sahukar, &koragappa.Koragappa, &tungamma.Tungamma, &akkamma.Akkamma, &choudappa.Chaudappa,
		},
	}
)
