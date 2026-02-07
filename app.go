package main

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"VideoCompressor/pkg/ffmpeg" // Import the new ffmpeg package
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

var bitrates = map[string]string{
	"1920x1080": "2500k",
	"1280x720":  "1000k",
	"854x480":   "500k",
}

func (a *App) Compress(inputPath, resolution string) (string, error) {
	bitrate, ok := bitrates[resolution]
	if !ok {
		return "", fmt.Errorf("Unknown resolution: %s", resolution)
	}
	ext := filepath.Ext(inputPath)
	outFile := strings.TrimSuffix(inputPath, ext) + "_compressed" + ext

	args := []string{
		"-i", inputPath,
		"-vf", fmt.Sprintf("scale=%s", resolution),
		"-b:v", bitrate,
		"-c:a", "copy",
		outFile,
	}

	stdout, stderr, err := ffmpeg.RunFFmpeg(args...)
	if err != nil {
		return "", fmt.Errorf("ffmpeg error: %v\nStdout: %s\nStderr: %s", err, stdout, stderr)
	}
	return outFile, nil
}

func (a *App) OpenFileDialog(ctx context.Context) (string, error) {
	result, err := wailsRuntime.OpenFileDialog(ctx, wailsRuntime.OpenDialogOptions{
		Title: "Select a video file",
		Filters: []wailsRuntime.FileFilter{
			{DisplayName: "Video files", Pattern: "*.mp4;*.mkv;*.avi"},
		},
	})
	return result, err
}
