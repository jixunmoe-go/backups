package main

import (
	"fmt"
	"github.com/jixunmoe-go/backups/utils/backup"
)

func printVerifyHelp() {
	println(appName + " verify [name] [timestamp]")
	println("")
	println("Verify one or multiple archives uploaded.")
	println("'name' and 'timestamp' are matched as prefix.")
	println("")
	println("e.g.")
	println("  " + appName + " verify")
	println("    Verify all uploaded archives checksum.")
	println("  " + appName + " verify db01")
	println("    Verify all archives for 'db01'.")
	println("  " + appName + " verify db01 1598918400")
	println("    Verify the archive uploaded to 'db01' in 2020.09.01")
}

func commandVerify(argv []string) int {
	name := ""
	time := ""
	if len(argv) > 1 {
		time = argv[1]
	}
	if len(argv) > 0 {
		name = argv[0]
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
	for _, project := range backup.GetProjectsWithPrefix(name) {
		for _, archive := range backup.GetBackupArchivesWithPrefix(project, time) {
			fmt.Printf(" [...] %s/%s\r", project, archive.FileName)
			verified, err := archive.Verify()
			printVerifyResult(verified)
			if err != nil {
				println(err)
			}
			if !verified {
				errors += 1
			}
		}
	}

	return errors
}
