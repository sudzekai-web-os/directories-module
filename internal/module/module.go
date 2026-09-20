package module

import (
	"github.com/sudzekai-web-os/core"
	"github.com/sudzekai-web-os/directories-module/internal/controllers/dirscontroller"
	"github.com/sudzekai-web-os/directories-module/internal/dispatchering/getfiles"
	"github.com/sudzekai-web-os/directories-module/internal/dispatchering/mkdir"
	"github.com/sudzekai-web-os/directories-module/internal/utilities/directoriessystem"
	"github.com/sudzekai-web-os/mediator"
)

type DirectoriesModule struct {
	registry      core.IHandlersRegistry
	loggerFactory core.ILoggerFactory
	executor      core.IExecutor
	configuration core.IConfiguration
}

func NewDirectoriesModule() *DirectoriesModule {
	return &DirectoriesModule{}
}

func (m *DirectoriesModule) Name() string {
	return "Directories Module"
}

func (m *DirectoriesModule) Description() string {
	return ""
}

func (m *DirectoriesModule) Version() string {
	return "v0.3.0"
}

func (m *DirectoriesModule) Start() error {
	dirSys := directoriessystem.NewDirectoriesSystem(m.executor)

	mediator.RegisterHandler(getfiles.NewHandler(dirSys))
	mediator.RegisterHandler(mkdir.NewHandler(dirSys))

	dirCtr := dirscontroller.New(m.loggerFactory)
	dirCtr.AddRoutes(m.registry)

	return nil
}

func (m *DirectoriesModule) AddHandlersRegistry(registry core.IHandlersRegistry) {
	m.registry = registry
}

func (m *DirectoriesModule) AddLoggerFactory(loggerFactory core.ILoggerFactory) {
	m.loggerFactory = loggerFactory
}

func (m *DirectoriesModule) AddExecutor(executor core.IExecutor) {
	m.executor = executor
}

func (m *DirectoriesModule) AddConfiguration(configuration core.IConfiguration) {
	m.configuration = configuration
}
