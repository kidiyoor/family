package mahabala

import (
  "kidiyoor.io/family-tree/pkg/types"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/mahabala/bhavani"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/mahabala/tilothama"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/mahabala/roopalatha"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/mahabala/padmaraja"
  "kidiyoor.io/family-tree/pkg/members/rangu/akkamma/mahabala/jayaraj"
)

var (
  // Mahabala is a member of the family.
  Mahabala = types.Member{
    Name: "Mahabala",
    Children: []*types.Member{&bhavani.Bhavani, &jayaraj.Jayaraj, &padmaraja.Padmaraja, &roopalatha.Roopalatha, &tilothama.Tilothama, },
    Gender: types.Male,
  }
)
