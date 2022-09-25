package sharath

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/sundari/sharath/sahithya"
)

var (
  // Sharath is a member of the family.
  Sharath = types.Member{
    Name: "Sharath",
    Children: []*types.Member{&sahithya.Sahithya, },
    Gender: types.Male,
  }
)
