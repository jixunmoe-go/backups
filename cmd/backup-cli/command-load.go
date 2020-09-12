package main

import (
	"flag"
	"github.com/jixunmoe-go/backups/utils/backup"
	"io"
	"os"
	"path"
	"strings"
)

func commandLoad(argv []string) int {
	command := flag.NewFlagSet("load", flag.ExitOnError)

	var projectName string
	var archiveTime string
	command.StringVar(&projectName, "name", "", "The name of the archive. e.g. blog-sql")
	command.StringVar(&archiveTime, "time", "latest", "The version of archive. Matches latest copy that contains this parameter as prefix.")
	if err := command.Parse(argv); err != nil {
		println("err: could not parse args: " + err.Error())
		command.PrintDefaults()
		return 2
	}

	if projectName == "" {
		println("err: -name is empty")
		command.PrintDefaults()
		return 1
	}

	return loadArchive(projectName, archiveTime)
}

func loadArchive(projectName, time string) int {
	var archive *backup.ArchiveInfo = nil
	archives := backup.GetBackupArchives(projectName)
	if time == "latest" {
		archive = archives[0]
	} else {
		for _, a := range archives {
			if strings.HasPrefix(a.FileName, time) {
				archive = a
				break
			}
		}
	}

	if archive == nil {
		println("could not find a valid archive. use the `list` command to see a list of valid times.")
		return 3
	}

	rootDir := backup.GetBackupLocation(projectName)
	f, err := os.OpenFile(path.Join(rootDir, archive.FileName), os.O_RDONLY, 0600)
	if err != nil {
		println("could not open backup.")
		return 4
	}

	_, _ = io.Copy(os.Stdout, f)
	_ = f.Close()

	return 0
}
