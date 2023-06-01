package screen

import (
	"os"
	"os/exec"
	"runtime"
)

// Clear will look for the current OS and clear the screen accordingly.
func Clear() {
	cmd := clearCmdBasedOnOS()
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		println("Warning: failed to clear screen:", err.Error())
	}
}

func clearCmdBasedOnOS() *exec.Cmd {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	return cmd
}
