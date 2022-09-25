package padhmalatha

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/harini/padhmalatha/sparsh"
)

var (
  // Padhmalatha is a member of the family.
  Padhmalatha = types.Member{
    Name: "Padmalatha",
    Spouse: &types.Member{
        Name: "Sandeep Kadur",
        Gender: types.Male,
      },
    Children: []*types.Member{&sparsh.Sparsh, },
    Gender: types.Female,
  }
)
