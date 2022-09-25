package yashoda

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/yashoda/madhukar"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/yashoda/dhanaraj"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/yashoda/mayur"
)

var (
  // Yashoda is a member of the family.
  Yashoda = types.Member{
    Name: "Yashoda",
    Children: []*types.Member{&dhanaraj.Dhanaraj, &madhukar.Madhukar, &mayur.Mayur, },
    Gender: types.Female,
  }
)
