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

