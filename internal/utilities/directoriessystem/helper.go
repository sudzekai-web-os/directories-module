package directoriessystem

import "path/filepath"

func IsAbsolute(path string) bool {
	return filepath.IsAbs(path)
}
