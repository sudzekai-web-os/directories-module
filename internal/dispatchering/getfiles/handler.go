package getfiles

import (
	"github.com/sudzekai-web-os/directories-module/internal/utilities/directoriessystem"
	"github.com/sudzekai-web-os/mediator"
)

type QueryHandler struct {
	dirSys *directoriessystem.DirectoriesSystem
}

func NewHandler(directoriesSystem *directoriessystem.DirectoriesSystem) mediator.IHandler[Query, QueryResult] {
	return QueryHandler{
		dirSys: directoriesSystem,
	}
}

func (h QueryHandler) Handle(q Query) QueryResult {
	files, err := h.dirSys.GetFiles(q.Path)

	return NewQueryResult(files, err)
}
