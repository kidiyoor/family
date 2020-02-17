package muthaakka

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/muthaakka/devdaspurushotham"
  "kidiyoor.io/family-tree/pkg/members/rangu/sheshasahukar/muthaakka/bhujangabanuvathi"
)

var (
  // Muthaakka is a member of the family.
  Muthaakka = types.Member{
    Name: "Muthaakka",
    Children: []*types.Member{&bhujangabanuvathi.Bhujangabanuvathi, &devdaspurushotham.Devdaspurushotham, },
    Gender: types.Female,
  }
)
