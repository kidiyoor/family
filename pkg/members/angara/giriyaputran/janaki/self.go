package janaki

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/janaki/mohan"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/janaki/vaman"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/janaki/renuka"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/janaki/bharathi"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/janaki/rekha"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/janaki/latha"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/janaki/harsha"
)

var (
  // Janaki is a member of the family.
  Janaki = types.Member{
    Name: "Janaki",
    Children: []*types.Member{&bharathi.Bharathi, &harsha.Harsha, &latha.Latha, &mohan.Mohan, &rekha.Rekha, &renuka.Renuka, &vaman.Vaman, },
    Gender: types.Female,
  }
)
