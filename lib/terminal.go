package lib

import "os"

// IsTerminal returns true if the file descriptor is a terminal
func IsTerminal(f *os.File) bool {
	// Check if file descriptor is a terminal
	fileInfo, err := f.Stat()
	if err != nil {
		return false
	}
	// CharDevice mode bit indicates a terminal
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}
