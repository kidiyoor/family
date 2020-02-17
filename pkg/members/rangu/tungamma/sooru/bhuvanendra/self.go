package bhuvanendra

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/bhuvanendra/yajnesh"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/bhuvanendra/brijesh"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/bhuvanendra/jithesh"
  "kidiyoor.io/family-tree/pkg/members/rangu/tungamma/sooru/bhuvanendra/bhavya"
)

var (
  // Bhuvanendra is a member of the family.
  Bhuvanendra = types.Member{
    Name: "Bhuvanendra",
    Children: []*types.Member{&bhavya.Bhavya, &brijesh.Brijesh, &jithesh.Jithesh, &yajnesh.Yajnesh, },
    Gender: types.Male,
  }
)
