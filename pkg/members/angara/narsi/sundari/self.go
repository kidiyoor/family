package sundari

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/vinod"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/hemalatha"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/shakunthala"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/praveen"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/pramila"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/sharath"
)

var (
  // Sundari is a member of the family.
  Sundari = types.Member{
    Name: "Sundari",
    Children: []*types.Member{&hemalatha.Hemalatha, &pramila.Pramila, &praveen.Praveen, &shakunthala.Shakunthala, &sharath.Sharath, &vinod.Vinod, },
    Gender: types.Female,
  }
)
