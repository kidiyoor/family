package sooru

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/tharanatha"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/vanaja"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/sharada"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/bhuvanendra"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/lalitha"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/kalavathi"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/hiriyanna"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/chitra"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/leela"
)

var (
  // Sooru is a member of the family.
  Sooru = types.Member{
    Name: "Sooru",
    Children: []*types.Member{&bhuvanendra.Bhuvanendra, &chitra.Chitra, &hiriyanna.Hiriyanna, &kalavathi.Kalavathi, &lalitha.Lalitha, &leela.Leela, &sharada.Sharada, &tharanatha.Tharanatha, &vanaja.Vanaja, },
    Gender: types.Female,
  }
)
