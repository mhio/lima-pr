//go:build !darwin && !windows

package hostagent

import (
	"context"

	"github.com/lima-vm/sshocker/pkg/ssh"
)

func forwardTCP(ctx context.Context, sshConfig *ssh.SSHConfig, port int, local, remote string, verb string) error {
	return forwardSSH(ctx, sshConfig, port, local, remote, verb, false)
}

func getFreeVSockPort() (int, error) {
	return 0, nil
}
