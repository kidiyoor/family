package harini

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/harini/harsha"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/harini/nithin"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/harini/padhmalatha"
)

var (
  // Harini is a member of the family.
  Harini = types.Member{
    Name: "Harini",
    Children: []*types.Member{&harsha.Harsha, &nithin.Nithin, &padhmalatha.Padhmalatha, },
    Gender: types.Female,
  }
)
