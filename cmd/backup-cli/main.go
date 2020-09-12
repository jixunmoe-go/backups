package main

import (
	"github.com/google/shlex"
	"os"
)

const appName = "backup-cli"

var fromSSHShell = false

var sshOnlyCommands = []string{
	"save",
	"load",
	"list",
	"verify",
}

func contains(list []string, handle string) bool {
	for _, s := range list {
		if s == handle {
			return true
		}
	}
	return false
}

func main() {
	os.Exit(handleCommand(os.Args))
}

func handleCommand(argv []string) int {
	if len(argv) == 1 {
		printHelp()
		return 1
	}

	command := argv[1]
	if fromSSHShell && !contains(sshOnlyCommands, command) {
		println("command '" + command + "' is not allowed over ssh.")
		return 1
	}

	switch command {
	case "ssh":
		return commandSSH()
	case "gen":
		return commandGenerate()
	case "pubkey":
		return commandPubKey()
	case "encrypt":
		return commandEncrypt(argv[2:])
	case "decrypt":
		return commandDecrypt(argv[2:])
	case "save":
		return commandSave(argv[2:])
	case "load":
		return commandLoad(argv[2:])
	case "list":
		return commandList(argv[2:])
	case "verify":
		return commandVerify(argv[2:])
	}

	println("Unknown command " + command + ".")
	println("")
	printHelp()

	return 9
}

func commandSSH() int {
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

	return handleCommand(append([]string{appName}, args...))
}

func printHelp() {
	println("usage: " + appName + " <command> [<args>]")
	println("Commands available: ")
	println("")
	println("SSH Execution")
	println(" ssh      Use '" + appName + " ssh' as forced command.")
	println("")
	println("Key related")
	println(" gen      Generate a private key.")
	println(" pubkey   Get public key from a given private key.")
	println("")
	println("Encryption")
	println(" encrypt  Encrypt bytes from stdin (with pubkey) and write to stdout.")
	println(" decrypt  Decrypt bytes from stdin (with privkey) and write to stdout.")
	println("")
	println("Backup management")
	println(" save     Save content received from stdin to a specified location.")
	println(" load     Load content stored in the backup server.")
	println(" list     List backup projects, or versions of a given backup project.")
	println(" verify   Verify all or a specific project/version.")
}
