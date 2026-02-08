package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"sync"

	"VideoCompressor/pkg/ffmpeg"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx               context.Context
	cancelCompression context.CancelFunc
	mu                sync.Mutex
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

func (a *App) Compress(inputPath, resolution, speed string) (string, error) {
	bitrate, ok := bitrates[resolution]
	if !ok {
		return "", fmt.Errorf("Unknown resolution: %s", resolution)
	}

	// Validate speed/preset
	// allowed: ultrafast, fast, medium. Default to medium if unknown.
	if speed != "ultrafast" && speed != "fast" && speed != "medium" {
		speed = "medium"
	}

	// Create a cancellable context for this compression task
	ctx, cancel := context.WithCancel(a.ctx)
	a.mu.Lock()
	a.cancelCompression = cancel
	a.mu.Unlock()

	defer func() {
		a.mu.Lock()
		if a.cancelCompression != nil {
			cancel() // Ensure cleanup
			a.cancelCompression = nil
		}
		a.mu.Unlock()
	}()

	ext := filepath.Ext(inputPath)
	outFile := strings.TrimSuffix(inputPath, ext) + "_compressed" + ext

	// 1. Get Duration using ffmpeg parsing
	durationSec, err := ffmpeg.GetDuration(ctx, inputPath)
	if err != nil {
		if ctx.Err() == context.Canceled {
			return "", fmt.Errorf("compression cancelled")
		}
		return "", fmt.Errorf("ffmpeg duration error: %v", err)
	}
	totalMicroseconds := int64(durationSec * 1000000)

	// 2. Run ffmpeg with progress
	args := []string{
		"-progress", "pipe:1",
		"-i", inputPath,
		"-vf", fmt.Sprintf("scale=%s", resolution),
		"-b:v", bitrate,
		"-preset", speed,
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

	stdout, stderr, err := ffmpeg.RunFFmpegProgress(ctx, onProgress, args...)
	if err != nil {
		// specific check for context canceled?
		if ctx.Err() == context.Canceled {
			os.Remove(outFile)
			return "", fmt.Errorf("compression cancelled")
		}
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

// beforeClose is called when the application is about to close
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.cancelCompression != nil {
		a.cancelCompression()
	}
	return false
}

func (a *App) Cancel() {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.cancelCompression != nil {
		a.cancelCompression()
	}
}

func (a *App) ShowInFileExplorer(path string) {
	cmd := exec.Command("open", "-R", path)
	cmd.Run()
}
