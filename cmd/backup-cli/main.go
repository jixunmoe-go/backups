package main

import (
	"fmt"
	"os"
)

const appName = "backup-cli"

func main() {
	initCommands()
	os.Exit(handleCommand(os.Args[1:]))
}

func handleCommand(argv []string) int {
	if len(argv) == 0 {
		printHelp()
		return 1
	}

	commandName := argv[0]
	commandArgs := argv[1:]
	command, ok := getCommand(commandName)
	if !ok {
		println(fmt.Sprintf("command '%s' is not supported", commandName))
		return 1
	}

	if fromSSHShell && !command.SSH {
		println(fmt.Sprintf("command '%s' is not allowed over ssh", commandName))
		return 1
	}

	if len(commandArgs) > 0 {
		if commandArgs[0] == "help" || commandArgs[0] == "?" {
			command.Help()
			return 0
		}
	}

	return command.Run(commandArgs)
}

func printHelp() {
	println("usage: " + appName + " <command> [<args>]")
	println("Commands available: ")

	for _, command := range commands {
		if command.Header != "" {
			println("")
			println(command.Header)
		}
		println(fmt.Sprintf(" %-10s %s", command.Name, command.Description))
	}
}
