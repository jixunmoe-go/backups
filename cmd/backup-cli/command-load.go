package main

import (
	"github.com/jixunmoe-go/backups/utils/backup"
	"io"
	"os"
	"path"
	"strings"
)

func printLoadHelp() {
	println(appName + " load <name> [time=latest]")
	println("")
	println("Load an existing backup and send to stdout.")
	println("")
	println("e.g.")
	println("  " + appName + " load blog-daily latest")
	println("    Loads the latest copy of 'blog-daily' backup.")
	println("  " + appName + " load blog-daily 1599609600")
	println("    Loads the local copy of 'blog-daily' backup, created in 2020-09-09.")
	println("(SSH Shell)")
	println(`  ssh backup@example.com load my-db-01 | backup-cli decrypt "$(cat private.key)" > my-db-01.tar.gz`)
	println("    Download and decrypt the archive, then save it to a file.")
}

func commandLoad(argv []string) int {
	if len(argv) == 0 {
		println("do not know which archive to load.")
		return 1
	}

	projectName := argv[0]
	archiveTime := ""
	if len(argv) > 1 {
		archiveTime = argv[1]
	}

	if projectName == "" {
		println("err: name is empty")
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
