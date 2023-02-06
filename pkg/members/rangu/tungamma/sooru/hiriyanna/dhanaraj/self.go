package dhanaraj

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/hiriyanna/dhanaraj/reyansh"
)

var (
  // Dhanaraj is a member of the family.
  Dhanaraj = types.Member{
    Name: "Dhanaraj & Shilpa",
    Children: []*types.Member{&reyansh.Reyansh, },
    Gender: types.Male,
  }
)
