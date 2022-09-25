package vinod

import (
"kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/prasanna/vineet"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/prasanna/suneet"  
"kidiyoor.io/family-tree/pkg/types"
)

var (
  // Vinod is a member of the family.
  Vinod = types.Member{
    Name: "Vinod",
    Children: []*types.Member{&suneet.Suneet, &vineet.Vineet, },
    Gender: types.Male,
  }
)
