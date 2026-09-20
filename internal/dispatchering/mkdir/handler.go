package mkdir

import (
	"github.com/sudzekai-web-os/directories-module/internal/utilities/directoriessystem"
	"github.com/sudzekai-web-os/mediator"
)

type CommandHandler struct {
	dirSys *directoriessystem.DirectoriesSystem
}

func NewHandler(directoriesSystem *directoriessystem.DirectoriesSystem) mediator.IHandler[Command, CommandResult] {
	return CommandHandler{
		dirSys: directoriesSystem,
	}
}

func (h CommandHandler) Handle(c Command) CommandResult {
	err := h.dirSys.CreateDirectory(c.Path, c.Name)

	return NewCommandResult(err)
}
