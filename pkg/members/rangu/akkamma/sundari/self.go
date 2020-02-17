package sundari

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/geetha"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/raychand"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/vagdevi"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/godavari"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/daskshayani"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/chandraprabha"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/preveena"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/sundari/premchand"
)

var (
  // Sundari is a member of the family.
  Sundari = types.Member{
    Name: "Sundari",
    Children: []*types.Member{&chandraprabha.Chandraprabha, &daskshayani.Daskshayani, &geetha.Geetha, &godavari.Godavari, &premchand.Premchand, &preveena.Preveena, &raychand.Raychand, &vagdevi.Vagdevi, },
    Gender: types.Female,
  }
)
