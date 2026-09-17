package module

import (
	"github.com/sudzekai-web-os/abstractions"
)

type DirectoriesModule struct {
}

func NewDirectoriesModule() abstractions.IModule {
	return &DirectoriesModule{}
}

func (m *DirectoriesModule) Name() string {
	return "Directories Module"
}

func (m *DirectoriesModule) Description() string {
	return ""
}

func (m *DirectoriesModule) Version() string {
	return "v0.9.0"
}

func (m *DirectoriesModule) Initialize(
	registry abstractions.IHandlersRegistry,
	loggerFactory abstractions.ILoggerFactory,
	executor abstractions.IExecutor,
) error {

	return nil
}
