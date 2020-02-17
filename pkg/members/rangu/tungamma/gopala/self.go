package gopala

import (
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/gopala/g"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/gopala/hemagiri"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/gopala/jayanthi"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/gopala/karunakara"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/gopala/shivdas"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/gopala/shrimathi"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/gopala/yuvaraj"
	"kidiyoor.io/family-tree/pkg/members/rangu/tungamma/gopala/yuvathikala"
	"kidiyoor.io/family-tree/pkg/types"
)

var (
	// Gopala is a member of the family.
	Gopala = types.Member{
		Name:     "Gopala",
		Children: []*types.Member{&g.G, &hemagiri.Hemagiri, &jayanthi.Jayanthi, &karunakara.Karunakara, &shivdas.Shivdas, &shrimathi.Shrimathi, &yuvaraj.Yuvaraj, &yuvathikala.Yuvathikala},
		Gender:   types.Male,
	}
)
