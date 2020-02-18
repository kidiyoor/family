package chaudappa

import (
	"kidiyoor.io/family-tree/pkg/members/rangu/chaudappa/achuthananda"
	"kidiyoor.io/family-tree/pkg/members/rangu/chaudappa/chandrashekara"
	"kidiyoor.io/family-tree/pkg/members/rangu/chaudappa/dayavathi"
	"kidiyoor.io/family-tree/pkg/members/rangu/chaudappa/devaki"
	"kidiyoor.io/family-tree/pkg/members/rangu/chaudappa/jahnavi"
	"kidiyoor.io/family-tree/pkg/members/rangu/chaudappa/purandara"
	"kidiyoor.io/family-tree/pkg/members/rangu/chaudappa/ramesha"
	"kidiyoor.io/family-tree/pkg/members/rangu/chaudappa/shivashankara"
	"kidiyoor.io/family-tree/pkg/types"
)

var (
	// Chaudappa is a member of the family.
	Chaudappa = types.Member{
		Name:     "Chandappa",
		Children: []*types.Member{&achuthananda.Achuthananda, &chandrashekara.Chandrashekara, &dayavathi.Dayavathi, &devaki.Devaki, &jahnavi.Jahnavi, &purandara.Purandara, &ramesha.Ramesha, &shivashankara.Shivashankara},
		Gender:   types.Male,
	}
)
