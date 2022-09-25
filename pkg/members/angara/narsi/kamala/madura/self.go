package madura

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/madura/gitesh"
  "kidiyoor.io/family-tree/pkg/members/angara/narsi/kamala/madura/daksh"
)

var (
  // Madura is a member of the family.
  Madura = types.Member{
    Name: "Madura",
    Children: []*types.Member{&daksh.Daksh, &gitesh.Gitesh, },
    Gender: types.Female,
  }
)
