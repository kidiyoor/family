package types

type Gender string

var Male Gender = "male"
var Female Gender = "female"

type Member struct {
	Name              string
	Gender          Gender
	Husband         *Member
	Wife            *Member
	Children        []*Member
	Photo           string
	Occupation      string
	Accomplishments []Accomplishment
  FacebookRef string
  TwitterRef string
}

type Accomplishment struct {
	Desc  string
	URL   string
	Photo string
}
