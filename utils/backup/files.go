package backup

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
)

var reFileName = regexp.MustCompile(`^(\d+)\.bin$`)

type ArchiveInfo struct {
	Name string
	Time uint64
}

func createArchiveFromName(name string) *ArchiveInfo {
	m := reFileName.FindStringSubmatch(name)
	ts, err := strconv.ParseUint(m[1], 10, 64)
	if err != nil {
		return nil
	}
	return &ArchiveInfo{
		Name: name,
		Time: ts,
	}
}

func GetBackupProjects() []string {
	root := GetBackupLocation(".")
	s, err := ioutil.ReadDir(root)
	if err != nil {
		return nil
	}
	var names []string
	for _, item := range s {
		if item.IsDir() {
			names = append(names, item.Name())
		}
	}
	return names
}

// GetBackupArchives returns a list of files (most recent ones at beginning)
func GetBackupArchives(name string) []*ArchiveInfo {
	var archives []*ArchiveInfo
	root := GetBackupLocation(name)
	s, err := ioutil.ReadDir(root)
	if err != nil {
		println("could not read dir")
		return nil
	}

	for _, v := range s {
		if !v.IsDir() {
			archive := createArchiveFromName(v.Name())

			if archive != nil {
				archives = append(archives, archive)
			}
		}
	}

	sort.SliceStable(archives, func(i, j int) bool {
		return archives[i].Time > archives[j].Time
	})

	return archives
}
