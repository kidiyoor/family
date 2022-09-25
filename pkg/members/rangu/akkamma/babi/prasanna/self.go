package prasanna

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/prasanna/vineet"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/prasanna/suneet"
)

var (
  // Prasanna is a member of the family.
  Prasanna = types.Member{
    Name: "Prasanna",
    Spouse: &types.Member{
        Name: "Vinod",
        Gender: types.Male,
      },
    Children: []*types.Member{&suneet.Suneet, &vineet.Vineet, },
    Gender: types.Female,
  }
)
