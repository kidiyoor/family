package vanaja

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/vanaja/vedu"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/vanaja/murali"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/vanaja/prakash"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/vanaja/uday"
  "kidiyoor.io/family-tree/pkg/members/angara/giriyaputran/vanaja/vaishali"
)

var (
  // Vanaja is a member of the family.
  Vanaja = types.Member{
    Name: "Vanaja",
    Children: []*types.Member{&murali.Murali, &prakash.Prakash, &uday.Uday, &vaishali.Vaishali, &vedu.Vedu, },
    Gender: types.Female,
  }
)
