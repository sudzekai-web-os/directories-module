package mkdir

type CommandResult struct {
	Error error
}

func NewCommandResult(err error) CommandResult {
	return CommandResult{
		Error: err,
	}
}
