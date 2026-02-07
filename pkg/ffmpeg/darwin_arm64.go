//go:build darwin && arm64

package ffmpeg

import (
	_ "embed"
)

//go:embed ffmpeg_binaries/darwin_arm64/ffmpeg
var ffmpegBinary []byte

//go:embed ffmpeg_binaries/darwin_arm64/ffprobe
var ffprobeBinary []byte

func runFFmpeg(args ...string) (string, string, error) {
	return runFFmpegInternal(ffmpegBinary, "ffmpeg", args...)
}

func runFFprobe(args ...string) (string, string, error) {
	return runFFmpegInternal(ffprobeBinary, "ffprobe", args...)
}

func runFFmpegProgress(onProgress func(string), args ...string) (string, string, error) {
	return runFFmpegProgressInternal(ffmpegBinary, "ffmpeg", onProgress, args...)
}
