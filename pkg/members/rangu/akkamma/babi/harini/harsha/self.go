package harsha

import (
	"kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/harini/harsha/akarsh"
)

var (
	// Harsha is a member of the family.
	Harsha = types.Member{
		Name:     "Harshalatha",
    Children: []*types.Member{&akarsh.Akarsh, },
		Gender:   types.Female,
	}
)
