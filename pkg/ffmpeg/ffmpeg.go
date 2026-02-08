package ffmpeg

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
)

// RunFFmpeg executes the ffmpeg binary with the given arguments.
// It returns stdout, stderr, and an error if the command fails.
// The actual implementation is provided by platform-specific files.
func RunFFmpeg(ctx context.Context, args ...string) (string, string, error) {
	return runFFmpeg(ctx, args...)
}

// GetDuration returns the duration of the media file in seconds.
func GetDuration(ctx context.Context, inputPath string) (float64, error) {
	// calling ffmpeg -i inputPath
	// expected to fail with exit code 1 because no output file is provided
	// however, stderr will contain the media information
	_, stderr, _ := RunFFmpeg(ctx, "-i", inputPath)

	// Regex to parse duration: Duration: 00:00:05.31
	re := regexp.MustCompile(`Duration: (\d{2}):(\d{2}):(\d{2}\.\d{2})`)
	matches := re.FindStringSubmatch(stderr)
	if len(matches) < 4 {
		return 0, fmt.Errorf("could not parse duration from ffmpeg output")
	}

	hours, _ := strconv.ParseFloat(matches[1], 64)
	minutes, _ := strconv.ParseFloat(matches[2], 64)
	seconds, _ := strconv.ParseFloat(matches[3], 64)

	return hours*3600 + minutes*60 + seconds, nil
}

// RunFFmpegProgress executes the ffmpeg binary with the given arguments and calls the onProgress callback for each line of output.
func RunFFmpegProgress(ctx context.Context, onProgress func(string), args ...string) (string, string, error) {
	return runFFmpegProgress(ctx, onProgress, args...)
}
