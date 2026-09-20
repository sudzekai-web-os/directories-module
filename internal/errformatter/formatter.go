package errformatter

import (
	"fmt"
	"net/http"

	"github.com/sudzekai-web-os/core"
	"github.com/sudzekai-web-os/directories-module/internal/utilities/directoriessystem"
)

var errors = map[error]core.HandlerResult{
	directoriessystem.ErrorAlreadyExists: {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("каталог уже существует"),
	},
	directoriessystem.ErrorAbsolutePath: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("путь должен быть абсолютным"),
	},
	directoriessystem.ErrorNotFound: {
		StatusCode: http.StatusNotFound,
		Error:      fmt.Errorf("указанный каталог не существует"),
	},
	directoriessystem.ErrorPermissionDenied: {
		StatusCode: http.StatusForbidden,
		Error:      fmt.Errorf("недостаточно прав для доступа к каталогу"),
	},
	directoriessystem.ErrorNotDirectory: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("указанный путь не является каталогом"),
	},
	directoriessystem.ErrorTooManySymlinks: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("слишком много уровней символических ссылок"),
	},
	directoriessystem.ErrorNameTooLong: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("имя каталога слишком длинное"),
	},
	directoriessystem.ErrorOutOfMemory: {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("не удалось выделить память"),
	},
	directoriessystem.ErrorIO: {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("ошибка ввода-вывода"),
	},
	directoriessystem.ErrorReadOnlyFileSystem: {
		StatusCode: http.StatusForbidden,
		Error:      fmt.Errorf("файловая система доступна только для чтения"),
	},
	directoriessystem.ErrorResourceBusy: {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("каталог или устройство занято"),
	},
	directoriessystem.ErrorNoSpace: {
		StatusCode: http.StatusInsufficientStorage,
		Error:      fmt.Errorf("недостаточно свободного места"),
	},
	directoriessystem.ErrorQuotaExceeded: {
		StatusCode: http.StatusInsufficientStorage,
		Error:      fmt.Errorf("превышена квота дискового пространства"),
	},
	directoriessystem.ErrorOperationNotPermitted: {
		StatusCode: http.StatusForbidden,
		Error:      fmt.Errorf("операция запрещена"),
	},
	directoriessystem.ErrorInvalidArgument: {
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf("передан некорректный аргумент"),
	},
	directoriessystem.ErrorInterrupted: {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("операция была прервана"),
	},
	directoriessystem.ErrorTooManyOpenFiles: {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("превышено количество открытых файлов"),
	},
	directoriessystem.ErrorDirectoryNotEmpty: {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("каталог не пуст"),
	},
	directoriessystem.ErrorTooManyLinks: {
		StatusCode: http.StatusConflict,
		Error:      fmt.Errorf("превышено количество ссылок на каталог"),
	},
	directoriessystem.ErrorNoSuchDevice: {
		StatusCode: http.StatusNotFound,
		Error:      fmt.Errorf("устройство не найдено"),
	},
	directoriessystem.ErrorNoSuchDeviceOrAddress: {
		StatusCode: http.StatusNotFound,
		Error:      fmt.Errorf("устройство или адрес не найден"),
	},
	directoriessystem.ErrorStaleFileHandle: {
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Errorf("файловый дескриптор устарел"),
	},
}

var defaultErr = core.HandlerResult{
	StatusCode: http.StatusInternalServerError,
	Error:      fmt.Errorf("неизвестная ошибка сервера"),
}

func MakeBusinessError(err error) core.HandlerResult {
	if result, exists := errors[err]; exists {
		return result
	}

	return defaultErr
}
