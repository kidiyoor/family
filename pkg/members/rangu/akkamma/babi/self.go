package babi

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/chandravathi"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/madhavi"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/dilipraj"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/lokanath"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/harini"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/venkatramana"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/devayani"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/babi/prasanna"
)

var (
  // Babi is a member of the family.
  Babi = types.Member{
    Name: "Babi",
    Children: []*types.Member{&chandravathi.Chandravathi, &devayani.Devayani, &dilipraj.Dilipraj, &harini.Harini, &lokanath.Lokanath, &madhavi.Madhavi, &prasanna.Prasanna, &venkatramana.Venkatramana, },
    Gender: types.Female,
  }
)
