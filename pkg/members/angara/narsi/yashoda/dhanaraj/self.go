package dhanaraj

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/yashoda/dhanaraj/dhanvi"
)

var (
  // Dhanaraj is a member of the family.
  Dhanaraj = types.Member{
    Name: "Dhanaraj",
    Children: []*types.Member{&dhanvi.Dhanvi, },
    Gender: types.Male,
  }
)
