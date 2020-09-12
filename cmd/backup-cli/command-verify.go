package main

import (
	"flag"
	"fmt"
	"github.com/jixunmoe-go/backups/utils/backup"
	"github.com/jixunmoe-go/backups/utils/checksum"
	"github.com/jixunmoe-go/backups/utils/dummy"
	"io"
	"os"
	"strings"
)

func commandVerify(argv []string) int {
	command := flag.NewFlagSet("verify", flag.ExitOnError)

	var name string
	var time string
	command.StringVar(&name, "name", "", "Project name prefix (colon separated, leave empty for all)")
	command.StringVar(&time, "time", "", "The timestamp prefix (colon separated, leave empty for all)")

	if err := command.Parse(argv); err != nil {
		println("err: could not parse args: " + err.Error())
		command.PrintDefaults()
		return 2
	}

	return verifyFiles(name, time)
}

func printVerifyResult(verified bool) {
	if verified {
		fmt.Printf(" [O K]\n")
	} else {
		fmt.Printf(" [BAD]\n")
	}
}

func verifyFiles(name, time string) int {
	errors := 0
	for _, project := range backup.GetBackupProjects() {
		if strings.HasPrefix(project, name) {
			for _, archive := range backup.GetBackupArchives(project) {
				if strings.HasPrefix(archive.FileName, time) {
					fmt.Printf(" [...] %s/%s\r", project, archive.FileName)
					f, err := os.OpenFile(archive.GetPath(), os.O_RDONLY, 0600)
					if err != nil {
						errors += 1
						printVerifyResult(false)
						break
					}
					reader := checksum.NewReader(f)
					_, _ = io.Copy(&dummy.Writer{}, reader)
					verified := reader.Verify()
					printVerifyResult(verified)
					if !verified {
						errors += 1
					}
				}
			}
		}
	}

	return errors
}
