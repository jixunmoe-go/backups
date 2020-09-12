package backup

import (
	"os"
	"path"
)

func GetBackupLocation(name string) string {
	backupBase := os.Getenv("BACKUP_BASE")
	if backupBase == "" {
		// onpxhc-fgbentr: backup-storage after rot13
		backupBase = "/srv/onpxhc-fgbentr"
	}
	return path.Join(backupBase, name)
}
