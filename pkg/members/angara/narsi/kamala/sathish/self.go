package sathish

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/sathish/vinyasa"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/sathish/vishruth"
)

var (
  // Sathish is a member of the family.
  Sathish = types.Member{
    Name: "Sathish",
    Children: []*types.Member{&vinyasa.Vinyasa, &vishruth.Vishruth, },
    Gender: types.Male,
  }
)
