package main

import (
	"fmt"
	"github.com/jixunmoe-go/backups/utils/backup"
	"time"
)

func printListHelp() {
	println(" " + appName + " list [name1] [name2] [...]")
	println("")
	println("List all backup names, or the version of specific backups.")
	println("Name is matched by prefix.")
	println("")
	println("e.g. ")
	println("  " + appName + " list      - list all backup names (project names)")
	println("  " + appName + " list test - list all backup timestamps available for 'test' projects")
}

func commandList(argv []string) int {
	if len(argv) < 1 {
		listProjects()
		return 0
	}

	for _, name := range argv {
		listArchive(name)
	}
	return 0
}

func listProjects() {
	projects := backup.GetBackupProjects()
	fmt.Printf("%d project(s) available\n", len(projects))
	for i, project := range projects {
		fmt.Printf(" %2d. %s\n", i+1, project)
	}
}

func listArchive(projectName string) {
	archives := backup.GetBackupArchives(projectName)
	fmt.Printf("%d version(s) available for %s\n", len(archives), projectName)

	for i, a := range archives {
		fmt.Printf(" %2d. %s (%s)\n", i+1, a.FileName, time.Unix(a.Time, 0).UTC().Format(time.UnixDate))
	}
}
