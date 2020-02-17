package hiriyanna

import (
	"kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/hiriyanna/suma"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/hiriyanna/dhanaraj"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/hiriyanna/prethiviraj"
)

var (
	// Hiriyanna is a member of the family.
	Hiriyanna = types.Member{
		Name:     "Hiriyanna",
    Children: []*types.Member{&dhanaraj.Dhanaraj, &prethiviraj.Prethiviraj, &suma.Suma, },
		Gender:   types.Male,
	}
)
