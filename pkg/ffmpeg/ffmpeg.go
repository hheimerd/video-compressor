package ffmpeg

// RunFFmpeg executes the ffmpeg binary with the given arguments.
// It returns stdout, stderr, and an error if the command fails.
// The actual implementation is provided by platform-specific files.
func RunFFmpeg(args ...string) (string, string, error) {
	return runFFmpeg(args...)
}

// RunFFprobe executes the ffprobe binary with the given arguments.
// It returns stdout, stderr, and an error if the command fails.
// The actual implementation is provided by platform-specific files.
func RunFFprobe(args ...string) (string, string, error) {
	return runFFprobe(args...)
}

// RunFFmpegProgress executes the ffmpeg binary with the given arguments and calls the onProgress callback for each line of output.
func RunFFmpegProgress(onProgress func(string), args ...string) (string, string, error) {
	return runFFmpegProgress(onProgress, args...)
}
