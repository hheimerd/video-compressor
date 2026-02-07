//go:build linux && amd64

package ffmpeg

import (
	_ "embed"
)

//go:embed ffmpeg_binaries/linux_amd64/ffmpeg
var ffmpegBinary []byte

//go:embed ffmpeg_binaries/linux_amd64/ffprobe
var ffprobeBinary []byte

func runFFmpeg(args ...string) (string, string, error) {
	return runFFmpegInternal(ffmpegBinary, "ffmpeg", args...)
}

func runFFprobe(args ...string) (string, string, error) {
	return runFFmpegInternal(ffprobeBinary, "ffprobe", args...)
}
