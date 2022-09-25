package krishnappakidiyoor

import (
"kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/dilipraj"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/lokanath"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/harini"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/venkatramana"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/devayani"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/prasanna"
  "kidiyoor.io/family-tree/pkg/types"
)

var (
  // Krishnappakidiyoor is a member of the family.
  Krishnappakidiyoor = types.Member{
    Name: "KrishnappaKidiyoor",
    Children: []*types.Member{&chandravathi.Chandravathi, &devayani.Devayani, &dilipraj.Dilipraj, &harini.Harini, &lokanath.Lokanath, &madhavi.Madhavi, &prasanna.Prasanna, &venkatramana.Venkatramana, },
    Gender: types.Male,
  }
)
