package shakunthala

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/shakunthala/vishwas"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/shakunthala/shreyas"
)

var (
  // Shakunthala is a member of the family.
  Shakunthala = types.Member{
    Name: "Shakunthala",
    Children: []*types.Member{&shreyas.Shreyas, &vishwas.Vishwas, },
    Gender: types.Female,
  }
)
