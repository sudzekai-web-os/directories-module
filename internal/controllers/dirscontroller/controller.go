package dirscontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sudzekai-web-os/core"
	"github.com/sudzekai-web-os/directories-module/internal/dispatchering/getfiles"
	"github.com/sudzekai-web-os/directories-module/internal/dispatchering/mkdir"
	"github.com/sudzekai-web-os/directories-module/internal/dto/requests"
	"github.com/sudzekai-web-os/directories-module/internal/errformatter"
	"github.com/sudzekai-web-os/mediator"
)

type Controller struct {
	loggerFactory core.ILoggerFactory
}

func New(loggerFactory core.ILoggerFactory) *Controller {
	return &Controller{
		loggerFactory: loggerFactory,
	}
}

func (ctr *Controller) AddRoutes(registry core.IHandlersRegistry) {
	registry.AddHandler("GET /directories/files", ctr.GetFiles)
	registry.AddHandler("POST /directories", ctr.CreateDirectory)
}

func (ctr *Controller) GetFiles(r *http.Request) core.HandlerResult {
	log := ctr.loggerFactory.NewLogger("controller-GET:directories/files")

	path := r.URL.Query().Get("path")

	var queryResult getfiles.QueryResult

	err := mediator.Dispatch(getfiles.NewQuery(path), &queryResult)

	if err == nil {
		err = queryResult.Error
	}
	if err != nil {
		log.LogError("ошибка получения файлов в директории: %s", err.Error())
		return errformatter.MakeBusinessError(err)
	}

	return core.HandlerResult{
		StatusCode: http.StatusOK,
		Data:       queryResult.FileInfos,
	}
}

func (ctr *Controller) CreateDirectory(r *http.Request) (result core.HandlerResult) {
	log := ctr.loggerFactory.NewLogger("controller-POST:directories/")

	path := r.URL.Query().Get("path")

	var body requests.DirectoryCreate

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		result.StatusCode = 400
		result.Error = fmt.Errorf("invalid json")
		return
	}

	var queryResult mkdir.CommandResult

	err := mediator.Dispatch(mkdir.NewCommand(path, body.Name), &queryResult)

	if err == nil {
		err = queryResult.Error
	}
	if err != nil {
		log.LogError("ошибка создания директории: %s", err.Error())
		return errformatter.MakeBusinessError(err)
	}

	return core.HandlerResult{
		StatusCode: http.StatusCreated,
	}
}
