package kalavathi

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/kalavathi/santosh"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/kalavathi/geetha"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/kalavathi/jaya"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/kalavathi/suma"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/kalavathi/dhanaraj"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/kalavathi/prethiviraj"
)

var (
  // Kalavathi is a member of the family.
  Kalavathi = types.Member{
    Name: "Kalavathi",
    Children: []*types.Member{&dhanaraj.Dhanaraj, &geetha.Geetha, &jaya.Jaya, &prethiviraj.Prethiviraj, &santosh.Santosh, &suma.Suma, },
    Gender: types.Female,
  }
)
