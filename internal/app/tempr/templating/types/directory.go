package templating

type Directory struct {
	Name		string		`yaml:"string"`
	Children	[]FSElement	`yaml:"children"`
}
