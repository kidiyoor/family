package kamala

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/somnath"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/ammi"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/sathish"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/mamtha"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/madura"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/shirish"
)

var (
  // Kamala is a member of the family.
  Kamala = types.Member{
    Name: "Kamala",
    Children: []*types.Member{&ammi.Ammi, &madura.Madura, &mamtha.Mamtha, &sathish.Sathish, &shirish.Shirish, &somnath.Somnath, },
    Gender: types.Female,
  }
)
