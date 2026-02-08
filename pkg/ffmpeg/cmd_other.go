//go:build !windows

package ffmpeg

import "os/exec"

func prepareCmd(cmd *exec.Cmd) {
	// No-op
}
