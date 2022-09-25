package giriyaputran

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/sanjeevbangera"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/rajubangera"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/devrajbangera"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/vanaja"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/janaki"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/leela"
)

var (
  // Giriyaputran is a member of the family.
  Giriyaputran = types.Member{
    Name: "GiriyaPutran",
    Children: []*types.Member{&devrajbangera.Devrajbangera, &janaki.Janaki, &leela.Leela, &rajubangera.Rajubangera, &sanjeevbangera.Sanjeevbangera, &vanaja.Vanaja, },
    Gender: types.Male,
  }
)
