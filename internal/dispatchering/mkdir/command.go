package mkdir

type Command struct {
	Path string
	Name string
}

func NewCommand(path, name string) Command {
	return Command{
		Path: path,
		Name: name,
	}
}
