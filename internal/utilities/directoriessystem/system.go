package directoriessystem

import (
	"path/filepath"
	"strings"

	"github.com/sudzekai-web-os/core"
)

type DirectoriesSystem struct {
	executor core.IExecutor
}

func NewDirectoriesSystem(executor core.IExecutor) *DirectoriesSystem {
	return &DirectoriesSystem{
		executor: executor,
	}
}

func (ds *DirectoriesSystem) CreateDirectory(dirPath, name string) error {
	if dir, _ := ds.IsDir(dirPath); dir {
		return ErrorAlreadyExists
	}

	cmdResult := ds.executor.Execute("mkdir", filepath.Join(dirPath, name))

	if err := formatCommandResult(cmdResult); err != nil {
		return err
	}

	return nil
}

func (ds *DirectoriesSystem) GetFiles(dirPath string) ([]*FileInfo, error) {
	result := make([]*FileInfo, 0)

	if dir, err := ds.IsDir(dirPath); !dir {
		return result, err
	}

	cmdResult := ds.executor.ExecuteInDirectory(
		dirPath,
		"sh",
		"-c",
		`for f in * .[^.]*; do [ -e "$f" ] && stat -c "N:%n|FN:$(realpath "$f")|MTIME:%Y|SIZE:%s|TYPE:%F" "$f"; done`,
	)

	if err := formatCommandResult(cmdResult); err != nil {
		return result, err
	}

	for line := range strings.SplitSeq(cmdResult.Stdout, "\n") {
		if line == "" {
			continue
		}

		result = append(result, NewFileInfo(line))
	}

	return result, nil
}

func (ds *DirectoriesSystem) IsDir(path string) (bool, error) {
	if !IsAbsolute(path) {
		return false, ErrorAbsolutePath
	}

	result := ds.executor.Execute("stat", "-c", "%F", path)

	if strings.TrimSpace(result.Stdout) != "directory" {
		return false, ErrorNotDirectory
	}

	return true, nil
}

func formatCommandResult(result core.CommandResult) error {
	if result.Stderr != "" {
		return ToError(result.Stderr)
	}

	if result.Error != nil {
		return ToError(result.Error.Error())
	}

	return nil
}
