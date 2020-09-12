package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(handleCommand(os.Args))
}

func handleCommand(argv []string) int {
	if len(argv) == 1 {
		printHelp()
		return 1
	}

	switch argv[1] {
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

	println("Unknown command " + argv[1] + ".")
	println("")
	printHelp()

	return 9
}

func printHelp() {
	println(fmt.Sprintf("usage: %s <command> [<args>]", os.Args[0]))
	println("Commands available: ")
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
	println(" verify   Verify all or a specific project/version, that is ")
}
