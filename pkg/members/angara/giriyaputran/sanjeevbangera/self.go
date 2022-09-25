package sanjeevbangera

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/sanjeevbangera/sandeep"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/sanjeevbangera/g"
)

var (
  // Sanjeevbangera is a member of the family.
  Sanjeevbangera = types.Member{
    Name: "SanjeevBangera",
    Children: []*types.Member{&g.G, &sandeep.Sandeep, },
    Gender: types.Male,
  }
)
