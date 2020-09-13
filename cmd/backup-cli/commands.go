package main

type Command struct {
	Header      string
	Name        string
	Description string
	Run         func(args []string) int
	Help        func()
	SSH         bool
}

var commands []*Command

func initCommands() {
	commands = []*Command{
		// SSH Shell
		{
			Header:      "SSH Shell",
			Run:         commandSSH,
			Help:        printSSHHelp,
			SSH:         false,
			Name:        "ssh",
			Description: "Enter SSH Shell mode and parse commands from $SSH_ORIGINAL_COMMAND.",
		},

		// Key related
		{
			Header:      "SSH Shell",
			Run:         commandGenerate,
			Help:        printGenerateHelp,
			SSH:         false,
			Name:        "gen",
			Description: "Generate a private key.",
		},
		{
			Run:         commandPubKey,
			Help:        printPubKeyHelp,
			SSH:         false,
			Name:        "pubkey",
			Description: "Get public key from a given private key.",
		},

		// Encryption
		{
			Header:      "Encryption",
			Run:         commandEncrypt,
			Help:        printEncryptHelp,
			SSH:         false,
			Name:        "encrypt",
			Description: "Encrypt bytes from stdin (with pubkey) and write to stdout.",
		},
		{
			Run:         commandDecrypt,
			Help:        printDecryptHelp,
			SSH:         false,
			Name:        "decrypt",
			Description: "Decrypt bytes from stdin (with privkey) and write to stdout.",
		},

		// Backup Management
		{
			Header:      "Backup Management",
			Run:         commandSave,
			Help:        printSaveHelp,
			SSH:         true,
			Name:        "save",
			Description: "Save content received from stdin to a specified location.",
		},
		{
			Run:         commandLoad,
			Help:        printLoadHelp,
			SSH:         true,
			Name:        "load",
			Description: "Load content stored in the backup server.",
		},
		{
			Run:         commandList,
			Help:        printListHelp,
			SSH:         true,
			Name:        "list",
			Description: "List backup projects, or versions of a given backup project.",
		},
		{
			Run:         commandVerify,
			Help:        printVerifyHelp,
			SSH:         true,
			Name:        "verify",
			Description: "Verify all or a specific project/version.",
		},
	}
}

func getCommand(name string) (cmd *Command, ok bool) {
	for _, cmd := range commands {
		if cmd.Name == name {
			return cmd, true
		}
	}
	return nil, false
}
