package getfiles

import "github.com/sudzekai-web-os/directories-module/internal/utilities/directoriessystem"

type QueryResult struct {
	FileInfos []*directoriessystem.FileInfo
	Error     error
}

func NewQueryResult(fileInfos []*directoriessystem.FileInfo, err error) QueryResult {
	return QueryResult{
		FileInfos: fileInfos,
		Error:     err,
	}
}
