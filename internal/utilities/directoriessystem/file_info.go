package directoriessystem

import (
	"strconv"
	"strings"
	"time"
)

type FileInfo struct {
	Name     string
	FullName string

	Type                 FileType
	Size                 int64
	ModificationDateTime time.Time

	stats map[string]string
}

// для корректной инициализации требуется следующий формат информации:
//
// N:apps|FN:/home/sudzekai/apps|MTIME:1789388874|SIZE:4096|TYPE:directory
//
// поиск осуществляется по ключевым словам, поэтому порядок не важен
//
// for f in * .[^.]*; do [ -e "$f" ] && stat -c "N:%n|FN:$(realpath "$f")|MTIME:%Y|SIZE:%s|TYPE:%F" "$f"; done
func NewFileInfo(stat string) *FileInfo {
	fileInfo := FileInfo{
		stats: make(map[string]string),
	}

	fileInfo.mapStat(stat)
	fileInfo.construct()

	return &fileInfo
}

func (fi *FileInfo) mapStat(stat string) {
	for entry := range strings.SplitSeq(stat, "|") {
		if idx := strings.Index(entry, ":"); idx != -1 {
			key := strings.TrimSpace(entry[:idx])
			value := strings.TrimSpace(entry[idx+1:])
			fi.stats[key] = value
		}
	}
}

func (fi *FileInfo) construct() {
	fi.Name = fi.getName()
	fi.FullName = fi.getFullName()

	fi.Type = fi.getType()

	fi.Size = fi.getSize()

	fi.ModificationDateTime = fi.getModificationDateTime()
}

func (fi *FileInfo) getName() string {
	return fi.stats["N"]
}

func (fi *FileInfo) getFullName() string {
	return fi.stats["FN"]
}

func (fi *FileInfo) getType() FileType {
	typeStr := fi.stats["TYPE"]

	if typeStr == "" {
		return TypeUnknown
	}

	switch typeStr {
	case "regular file", "regular empty file":
		return TypeRegular
	case "directory":
		return TypeDirectory
	case "symbolic link":
		return TypeSymlink
	case "socket":
		return TypeSocket
	case "fifo":
		return TypePipe
	case "block special file":
		return TypeBlockDevice
	case "character special file":
		return TypeCharacterDevice
	default:
		return TypeUnknown
	}
}

func (fi *FileInfo) getSize() int64 {
	sizeStr := fi.stats["SIZE"]
	size, _ := strconv.ParseInt(sizeStr, 10, 64)
	return size
}

func (fi *FileInfo) getModificationDateTime() time.Time {
	return parseUnixString(fi.stats["MTIME"])
}

func parseUnixString(unixStr string) time.Time {
	timestamp, _ := strconv.ParseInt(unixStr, 10, 64)

	t := time.Unix(timestamp, 0)
	return t
}
