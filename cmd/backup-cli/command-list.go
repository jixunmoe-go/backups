package main

import (
	"flag"
	"fmt"
	"github.com/jixunmoe-go/backups/utils/backup"
)

func commandList(argv []string) int {
	command := flag.NewFlagSet("load", flag.ExitOnError)

	var projectName string
	command.StringVar(&projectName, "name", "", "The name of the archive. e.g. blog-sql. omit to list names")
	if err := command.Parse(argv); err != nil {
		println("err: could not parse args: " + err.Error())
		command.PrintDefaults()
		return 2
	}

	if projectName == "" {
		return listProjects()
	}

	return listArchive(projectName)
}

func listProjects() int {
	projects := backup.GetBackupProjects()
	fmt.Printf("%d project(s) available\n", len(projects))
	for i, project := range projects {
		fmt.Printf(" %2d. %s\n", i+1, project)
	}
	return 0
}

func listArchive(projectName string) int {
	archives := backup.GetBackupArchives(projectName)
	fmt.Printf("%d version(s) available\n", len(archives))

	for i, a := range archives {
		fmt.Printf(" %2d. %s\n", i+1, a.FileName)
	}
	return 0
}
