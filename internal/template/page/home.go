package page

type Home struct {
	Text string
}

func (h Home) Name() string {
	return "home.gohtml"
}
