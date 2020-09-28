package main

import (
	"fmt"
	"github.com/jixunmoe-go/backups/utils/backup"
	"strconv"
)

func printCleanHelp() {
	println(appName + " clean [n=5]")
	println(appName + " clean [n=5] [name1] [name2] ...")
	println("")
	println("Remove old backup archives. Keep the latest n copies.")
	println("")
	println("e.g.")
	println("  " + appName + " clean")
	println("  " + appName + " clean 5 blog-daily")
}

func commandClean(args []string) int {
	var err error
	var projects []string

	// Default to keep 5
	keepCount := 5
	if len(args) > 0 {
		keepCount, err = strconv.Atoi(args[0])
		if err != nil {
			println("could not parse [n]: " + err.Error())
			return 1
		}
		projects = args[1:]
	}

	if len(projects) == 0 {
		projects = backup.GetBackupProjects()
	}

	bad := 0
	for _, project := range projects {
		ok := 0
		for _, archive := range backup.GetBackupArchives(project) {
			fmt.Printf(" [...] %s/%s\r", project, archive.FileName)
			if ok >= keepCount {
				fmt.Printf(" [DEL]\n")
				archive.Delete()
				continue
			}

			success, _ := archive.Verify()
			if success {
				fmt.Printf(" [O K]\n")
				ok++
			} else {
				fmt.Printf(" [VEF]\n")
				archive.Delete()
				bad++
			}
		}
	}
	return bad
}
