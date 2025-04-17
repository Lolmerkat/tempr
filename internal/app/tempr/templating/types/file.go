package templating

type File struct {
	Name		string		`yaml:"name"`
	Content		[]string	`yaml:"content"`
}
