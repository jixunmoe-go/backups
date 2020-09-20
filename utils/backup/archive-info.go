package backup

import (
	"fmt"
	"github.com/jixunmoe-go/backups/utils/humanize"
	"os"
	"path"
	"regexp"
	"strconv"
)

var reFileName = regexp.MustCompile(`^(\d+)\.bin$`)

type ArchiveInfo struct {
	FileName string
	Time     int64
	Project  string
}

func createArchiveFromName(project, name string) *ArchiveInfo {
	m := reFileName.FindStringSubmatch(name)
	ts, err := strconv.ParseInt(m[1], 10, 64)
	if err != nil {
		return nil
	}
	return &ArchiveInfo{
		FileName: name,
		Time:     ts,
		Project:  project,
	}
}

func (a *ArchiveInfo) GetPath() string {
	return path.Join(GetBackupLocation(a.Project), a.FileName)
}

func (a *ArchiveInfo) GetFormattedSize() string {
	fi, err := os.Stat(a.GetPath())
	if err != nil {
		return fmt.Sprintf("(unknown size: %s)", err)
	}
	return humanize.ByteCountBinary(fi.Size())
}
