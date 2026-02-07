package ffmpeg

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func runFFmpegInternal(binaryData []byte, executableName string, args ...string) (string, string, error) {
	tempDir, err := ioutil.TempDir("", "ffmpeg-")
	if err != nil {
		return "", "", err
	}
	defer os.RemoveAll(tempDir)

	ffmpegPath := filepath.Join(tempDir, executableName)

	// Set permissions to 0755 (rwx-rx-rx) to ensure the binary is executable
	// On Windows, this is not strictly necessary for .exe files but is good practice.
	err = ioutil.WriteFile(ffmpegPath, binaryData, 0755)
	if err != nil {
		return "", "", err
	}

	cmd := exec.Command(ffmpegPath, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return stdout.String(), stderr.String(), err
	}

	return stdout.String(), stderr.String(), nil
}
