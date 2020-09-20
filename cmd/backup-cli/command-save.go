package main

import (
	"fmt"
	"github.com/jixunmoe-go/backups/utils/backup"
	"io"
	"os"
	"path"
	"time"
)

func printSaveHelp() {
	println(appName + " save <name>")
	println("")
	println("Save data received from stdin pipe to a pre-defined backup storage location.")
	println("")
	println("e.g.")
	println("  cat backup.tar.gz | " + appName + " save my-db-01")
	println("    Save a backup file to backup storage.")
	println("(SSH Shell)")
	println("  cat backup.tar.gz | ssh backup@example.com save my-db-01")
}

func commandSave(argv []string) int {
	if len(argv) == 0 {
		println("missing argument <name>")
		return 1
	}

	projectName := argv[0]

	if projectName == "" {
		println("err: <name> is empty")
		return 1
	}

	return copyToFile(projectName)
}

func copyToFile(projectName string) int {
	backupPath := backup.GetBackupLocation(projectName)
	err := os.MkdirAll(backupPath, 0700)
	if err != nil {
		println("err: could not create backup dir: " + err.Error())
		return 1
	}

	backupArchive := path.Join(backupPath, fmt.Sprintf("%d.bin", time.Now().UTC().Unix()))
	f, err := os.OpenFile(backupArchive, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		println("err: could not create file: " + err.Error())
		return 1
	}

	_, err = io.Copy(f, os.Stdin)
	if err != nil {
		_ = f.Close()
		_ = os.Remove(backupArchive)
		return 9
	}
	_ = f.Close()
	return 0
}
