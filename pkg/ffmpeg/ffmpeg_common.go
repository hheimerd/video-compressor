package ffmpeg

import (
	"bufio"
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
)

func runFFmpegInternal(ctx context.Context, binaryData []byte, executableName string, args ...string) (string, string, error) {
	tempDir, err := os.MkdirTemp("", "ffmpeg-")
	if err != nil {
		return "", "", err
	}
	defer os.RemoveAll(tempDir)

	ffmpegPath := filepath.Join(tempDir, executableName)

	// Set permissions to 0755 (rwx-rx-rx) to ensure the binary is executable
	// On Windows, this is not strictly necessary for .exe files but is good practice.
	err = os.WriteFile(ffmpegPath, binaryData, 0755)
	if err != nil {
		return "", "", err
	}

	cmd := exec.CommandContext(ctx, ffmpegPath, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return stdout.String(), stderr.String(), err
	}

	return stdout.String(), stderr.String(), nil
}

func runFFmpegProgressInternal(ctx context.Context, binaryData []byte, executableName string, onProgress func(string), args ...string) (string, string, error) {
	tempDir, err := os.MkdirTemp("", "ffmpeg-")
	if err != nil {
		return "", "", err
	}
	defer os.RemoveAll(tempDir)

	ffmpegPath := filepath.Join(tempDir, executableName)

	// Set permissions to 0755 (rwx-rx-rx) to ensure the binary is executable
	// On Windows, this is not strictly necessary for .exe files but is good practice.
	err = os.WriteFile(ffmpegPath, binaryData, 0755)
	if err != nil {
		return "", "", err
	}

	cmd := exec.CommandContext(ctx, ffmpegPath, args...)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", "", err
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return "", stderr.String(), err
	}

	var stdoutBuf bytes.Buffer

	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		line := scanner.Text()
		stdoutBuf.WriteString(line + "\n")
		// FFmpeg with -progress pipe:1 writes key=value pairs.
		// We pass every line to the callback.
		onProgress(line)
	}

	if err := cmd.Wait(); err != nil {
		return stdoutBuf.String(), stderr.String(), err
	}

	return stdoutBuf.String(), stderr.String(), nil
}
