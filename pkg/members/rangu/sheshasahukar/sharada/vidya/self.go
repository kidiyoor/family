package vidya

import (
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Vidya is a member of the family.
  Vidya = types.Member{
    Name: "vidya",
    Children: []*types.Member{},
    Gender: types.Female,
  }
)
