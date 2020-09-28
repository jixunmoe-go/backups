package backup

import (
	"fmt"
	"github.com/jixunmoe-go/backups/utils/checksum"
	"github.com/jixunmoe-go/backups/utils/dummy"
	"github.com/jixunmoe-go/backups/utils/humanize"
	"io"
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

func (a *ArchiveInfo) Verify() (bool, error) {
	f, err := os.OpenFile(a.GetPath(), os.O_RDONLY, 0600)
	if err != nil {
		return false, err
	}
	reader := checksum.NewReader(f)
	_, _ = io.Copy(&dummy.Writer{}, reader)
	return reader.Verify(), nil
}

func (a *ArchiveInfo) Delete() {
	_ = os.Remove(a.GetPath())
}
