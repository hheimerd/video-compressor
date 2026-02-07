package main

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"VideoCompressor/pkg/ffmpeg"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
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

	// 1. Get Duration using ffprobe
	probeArgs := []string{
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		inputPath,
	}
	durationStr, probeErrStr, err := ffmpeg.RunFFprobe(probeArgs...)
	if err != nil {
		return "", fmt.Errorf("ffprobe error: %v\nstderr: %s", err, probeErrStr)
	}

	durationSec, _ := strconv.ParseFloat(strings.TrimSpace(durationStr), 64)
	totalMicroseconds := int64(durationSec * 1000000)

	// 2. Run ffmpeg with progress
	args := []string{
		"-progress", "pipe:1",
		"-i", inputPath,
		"-vf", fmt.Sprintf("scale=%s", resolution),
		"-b:v", bitrate,
		"-c:a", "copy",
		"-y", // Overwrite output file
		outFile,
	}

	onProgress := func(line string) {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 && strings.TrimSpace(parts[0]) == "out_time_us" {
			outTimeUs, _ := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
			if totalMicroseconds > 0 {
				percent := int(float64(outTimeUs) / float64(totalMicroseconds) * 100)
				if percent > 100 {
					percent = 100
				}
				wailsRuntime.EventsEmit(a.ctx, "progress", percent)
			}
		}
	}

	stdout, stderr, err := ffmpeg.RunFFmpegProgress(onProgress, args...)
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
