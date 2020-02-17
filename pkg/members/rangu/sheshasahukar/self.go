package sheshasahukar

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/neelamma"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/muthaakka"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/kamala"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/nagaveni"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/sharada"
)

var (
  // Sheshasahukar is a member of the family.
  Sheshasahukar = types.Member{
    Name: "Shesha sahukar",
    Children: []*types.Member{&kamala.Kamala, &muthaakka.Muthaakka, &nagaveni.Nagaveni, &neelamma.Neelamma, &sharada.Sharada, },
    Gender: types.Male,
  }
)
