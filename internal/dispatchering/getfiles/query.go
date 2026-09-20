package getfiles

type Query struct {
	Path string
}

func NewQuery(path string) Query {
	return Query{
		Path: path,
	}
}
