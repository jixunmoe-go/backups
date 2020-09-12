package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

func commandSave(argv []string) int {
	command := flag.NewFlagSet("save", flag.ExitOnError)

	var projectName string
	command.StringVar(&projectName, "name", "", "The name of the archive. e.g. blog-sql")
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

	return copyToFile(projectName)
}

func copyToFile(projectName string) int {
	backupBase := os.Getenv("BACKUP_BASE")
	if backupBase == "" {
		// onpxhc-fgbentr: backup-storage after rot13
		backupBase = "/srv/onpxhc-fgbentr"
	}
	backupPath := path.Join(backupBase, projectName)
	err := os.MkdirAll(backupPath, 0700)
	if err != nil {
		println("err: could not create backup dir: " + err.Error())
		return 1
	}

	backupArchive := path.Join(backupPath, fmt.Sprintf("%d.bin", time.Now().Unix()))
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
