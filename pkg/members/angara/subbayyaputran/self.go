package subbayyaputran

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran/vittal"
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran/venkatesh"
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran/raju"
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran/sarojini"
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran/pushpakargi"
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran/rama"
  "kidiyoor.io/family-tree/pkg/members/angara/subbayyaputran/kamal"
)

var (
  // Subbayyaputran is a member of the family.
  Subbayyaputran = types.Member{
    Name: "SubbayyaPutran",
    Children: []*types.Member{&kamal.Kamal, &pushpakargi.Pushpakargi, &raju.Raju, &rama.Rama, &sarojini.Sarojini, &venkatesh.Venkatesh, &vittal.Vittal, },
    Gender: types.Male,
  }
)
