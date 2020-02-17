package tungamma

import (
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ananda"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/annayya"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/gopala"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/guruva"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/ramappa"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru"
	"kidiyoor.io/family-tree/pkg/types"
)

var (
	// Tungamma is a member of the family.
	Tungamma = types.Member{
		Name:     "Tungamma",
		Children: []*types.Member{&ananda.Ananda, &annayya.Annayya, &gopala.Gopala, &guruva.Guruva, &ramappa.Ramappa, &sooru.Sooru},
		Gender:   types.Female,
	}
)
