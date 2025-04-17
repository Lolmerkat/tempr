package templating

type Template struct {
	Name		string		`yaml:"name"`
	Author		string		`yaml:"author,omitempty"`
	Version		string		`yaml:"version,omitempty"`
	Languages	[]string	`yaml:"languages,omitempty"`
	Content		[]FSElement	`yaml:"content"`
}
