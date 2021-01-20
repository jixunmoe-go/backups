package backup

import (
	"io/ioutil"
	"sort"
	"strings"
)

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
func GetBackupArchives(project string) []*ArchiveInfo {
	var archives []*ArchiveInfo
	root := GetBackupLocation(project)
	dirContent, err := ioutil.ReadDir(root)
	if err != nil {
		println("could not read dir")
		return nil
	}

	for _, file := range dirContent {
		if !file.IsDir() {
			archive := createArchiveFromName(project, file.Name())

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

func GetProjectsWithPrefix(prefix string) []string {
	var names []string
	for _, project := range GetBackupProjects() {
		if strings.HasPrefix(project, prefix) {
			names = append(names, project)
		}
	}
	return names
}

func GetBackupArchivesWithPrefix(project, prefix string) []*ArchiveInfo {
	var archives []*ArchiveInfo
	for _, archive := range GetBackupArchives(project) {
		if strings.HasPrefix(archive.FileName, prefix) {
			archives = append(archives, archive)
		}
	}
	return archives
}
