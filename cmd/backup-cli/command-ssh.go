package main

import (
	"github.com/google/shlex"
	"os"
)

var fromSSHShell = false

func printSSHHelp() {
	println(appName + " ssh")
	println("")
	println("Execute a restricted subset of commands via $SSH_ORIGINAL_COMMAND.")
	println("See README for help to setup `backup-cli` as the SSH Shell.")
}

func commandSSH(_ []string) int {
	fromSSHShell = true

	// Safe SSH command line parsing!
	sshCommand := os.Getenv("SSH_ORIGINAL_COMMAND")
	if sshCommand == "" {
		println("$SSH_ORIGINAL_COMMAND is empty.")
		return 1
	}

	args, err := shlex.Split(sshCommand)
	if err != nil {
		println("could not parse $SSH_ORIGINAL_COMMAND: " + err.Error())
		return 2
	}

	return handleCommand(args)
}
