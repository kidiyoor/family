package types

// Gender M/F.
type Gender string

// Gender Male.
var Male Gender = "male"

// Gender Female.
var Female Gender = "female"

// Member stores details of a family member.
type Member struct {
	Name            string `json:"name"`
	Gender          Gender
	Husband         *Member
	Wife            *Member
	Children        []*Member `json:"children"`
	Photo           string
	Occupation      string
	Accomplishments []Accomplishment
	FacebookRef     string
	TwitterRef      string
}

// Accomplishment of family members.
type Accomplishment struct {
	Desc  string
	URL   string
	Photo string
}
