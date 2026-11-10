// Package notify runs a user-supplied command when an interesting line shows up.
package notify

import "os/exec"

// Run executes cmd with the log line as its argument.
func Run(cmd, line string) error {
	return exec.Command("sh", "-c", cmd+" '"+line+"'").Run()
}
