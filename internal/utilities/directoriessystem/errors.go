package directoriessystem

import (
	"fmt"
	"strings"
)

var (
	ErrorAlreadyExists         = fmt.Errorf("Directory already exists")
	ErrorAbsolutePath          = fmt.Errorf("Path must be absolute")
	ErrorNotFound              = fmt.Errorf("Directory not found")
	ErrorPermissionDenied      = fmt.Errorf("Permission denied")
	ErrorNotDirectory          = fmt.Errorf("Not a directory")
	ErrorTooManySymlinks       = fmt.Errorf("Too many levels of symbolic links")
	ErrorNameTooLong           = fmt.Errorf("File name too long")
	ErrorOutOfMemory           = fmt.Errorf("Cannot allocate memory")
	ErrorIO                    = fmt.Errorf("Input/output error")
	ErrorReadOnlyFileSystem    = fmt.Errorf("Read-only file system")
	ErrorResourceBusy          = fmt.Errorf("Device or resource busy")
	ErrorNoSpace               = fmt.Errorf("No space left on device")
	ErrorQuotaExceeded         = fmt.Errorf("Disk quota exceeded")
	ErrorOperationNotPermitted = fmt.Errorf("Operation not permitted")
	ErrorInvalidArgument       = fmt.Errorf("Invalid argument")
	ErrorInterrupted           = fmt.Errorf("Interrupted system call")
	ErrorTooManyOpenFiles      = fmt.Errorf("Too many open files")
	ErrorDirectoryNotEmpty     = fmt.Errorf("Directory not empty")
	ErrorTooManyLinks          = fmt.Errorf("Too many links")
	ErrorNoSuchDevice          = fmt.Errorf("No such device")
	ErrorNoSuchDeviceOrAddress = fmt.Errorf("No such device or address")
	ErrorStaleFileHandle       = fmt.Errorf("Stale file handle")
)

var errors = map[string]error{
	"No such file or directory":         ErrorNotFound,
	"Permission denied":                 ErrorPermissionDenied,
	"Not a directory":                   ErrorNotDirectory,
	"Too many levels of symbolic links": ErrorTooManySymlinks,
	"File name too long":                ErrorNameTooLong,
	"Cannot allocate memory":            ErrorOutOfMemory,
	"Input/output error":                ErrorIO,
	"Read-only file system":             ErrorReadOnlyFileSystem,
	"File system is read-only":          ErrorReadOnlyFileSystem,
	"Device or resource busy":           ErrorResourceBusy,
	"No space left on device":           ErrorNoSpace,
	"Disk quota exceeded":               ErrorQuotaExceeded,
	"Operation not permitted":           ErrorOperationNotPermitted,
	"Invalid argument":                  ErrorInvalidArgument,
	"Interrupted system call":           ErrorInterrupted,
	"Too many open files":               ErrorTooManyOpenFiles,
	"Directory not empty":               ErrorDirectoryNotEmpty,
	"Too many links":                    ErrorTooManyLinks,
	"No such device":                    ErrorNoSuchDevice,
	"No such device or address":         ErrorNoSuchDeviceOrAddress,
	"Stale file handle":                 ErrorStaleFileHandle,
	"File exists":                       ErrorAlreadyExists,
}

func ToError(stdErr string) error {
	for stdErrSubStr, err := range errors {
		if strings.Contains(stdErr, stdErrSubStr) {
			return err
		}
	}

	return fmt.Errorf("%s", stdErr)
}
