//go:build windows && amd64

package ffmpeg

import (
	_ "embed"
)

//go:embed ffmpeg_binaries/windows_amd64/ffmpeg.exe
var ffmpegBinary []byte

//go:embed ffmpeg_binaries/windows_amd64/ffprobe.exe
var ffprobeBinary []byte

func runFFmpeg(args ...string) (string, string, error) {
	return runFFmpegInternal(ffmpegBinary, "ffmpeg.exe", args...)
}

func runFFprobe(args ...string) (string, string, error) {
	return runFFmpegInternal(ffprobeBinary, "ffprobe.exe", args...)
}
