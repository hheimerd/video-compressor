package ffmpeg

import "context"

// RunFFmpeg executes the ffmpeg binary with the given arguments.
// It returns stdout, stderr, and an error if the command fails.
// The actual implementation is provided by platform-specific files.
func RunFFmpeg(ctx context.Context, args ...string) (string, string, error) {
	return runFFmpeg(ctx, args...)
}

// RunFFprobe executes the ffprobe binary with the given arguments.
// It returns stdout, stderr, and an error if the command fails.
// The actual implementation is provided by platform-specific files.
func RunFFprobe(ctx context.Context, args ...string) (string, string, error) {
	return runFFprobe(ctx, args...)
}

// RunFFmpegProgress executes the ffmpeg binary with the given arguments and calls the onProgress callback for each line of output.
func RunFFmpegProgress(ctx context.Context, onProgress func(string), args ...string) (string, string, error) {
	return runFFmpegProgress(ctx, onProgress, args...)
}
