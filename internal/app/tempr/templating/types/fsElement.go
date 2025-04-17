package templating

type FSElement interface {
	Expand(path string)
}

type RawFSElement struct {
	Name		string			`yaml:"name"`
	Content		[]string		`yaml:"content"`
	Children	[]RawFSElement	`yaml:"children,omitempty"`
}
