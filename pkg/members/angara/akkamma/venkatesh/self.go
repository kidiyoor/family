package venkatesh

import (
"kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi/bipin"
	"kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi/dinnu"
	"kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi/sachin"
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Venkatesh is a member of the family.
  Venkatesh = types.Member{
    Name: "Venkatesh",
    Children: []*types.Member{&bipin.Bipin, &dinnu.Dinnu, &sachin.Sachin},
    Gender: types.Male,
  }
)
