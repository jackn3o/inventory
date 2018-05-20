// +build linux,amd64 darwin

package app

import (
	"os"
	"syscall"
)

var signals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGKILL,
	syscall.SIGTERM,
	syscall.SIGSTOP,
	syscall.SIGHUP,
	syscall.SIGQUIT,
}
