package jagadvira

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/girija/jagadvira/mamtha"
)

var (
  // Jagadvira is a member of the family.
  Jagadvira = types.Member{
    Name: "Jagadvira",
    Children: []*types.Member{&mamtha.Mamtha, },
    Gender: types.Male,
  }
)
