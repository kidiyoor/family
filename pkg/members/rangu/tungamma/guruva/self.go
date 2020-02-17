package guruva

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/guruva/vasantha"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/guruva/varija"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/guruva/malathi"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/guruva/annaji"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/guruva/damodara"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/guruva/jagadish"
)

var (
  // Guruva is a member of the family.
  Guruva = types.Member{
    Name: "Guruva",
    Children: []*types.Member{&annaji.Annaji, &damodara.Damodara, &jagadish.Jagadish, &malathi.Malathi, &varija.Varija, &vasantha.Vasantha, },
    Gender: types.Male,
  }
)
