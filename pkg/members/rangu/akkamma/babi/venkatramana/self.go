package venkatramana

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/venkatramana/gautham"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/venkatramana/dhanaraj"
)

var (
  // Venkatramana is a member of the family.
  Venkatramana = types.Member{
    Name: "Venkatramana",
    Children: []*types.Member{&dhanaraj.Dhanaraj, &gautham.Gautham, },
    Gender: types.Male,
  }
)
