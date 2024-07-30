package featnix

import (
	"os"
	"os/exec"
)

// HasSystemd checks if the systemd service manager is available.
//
// It checks if systemd is installed by checking for the presence of certain directories
// and the 'systemctl' command in the PATH.
//
// Returns:
//
//	bool: True if systemd service manager is available, false otherwise.
func HasSystemd() bool {
	// Directories where systemd stores its state files, unit files, and drop-in unit files.
	// These directories are checked to determine if systemd is installed.
	dirs := []string{
		"/run/systemd/system",     // State files.
		"/usr/lib/systemd/system", // Unit files.
		"/etc/systemd/system",     // Drop-in unit files.
	}

	// Check if any of the directories exist.
	for _, dir := range dirs {
		// Check if the directory exists.
		// If the directory exists, it indicates that systemd is installed.
		if _, err := os.Stat(dir); err == nil {
			return true
		}
	}

	// Check if the 'systemctl' command is in the PATH.
	// If the command is in the PATH, it indicates that systemd is installed.
	_, err := exec.LookPath("systemctl")

	// If there is no error, systemctl command is found in the PATH and systemd is installed.
	return err == nil
}
