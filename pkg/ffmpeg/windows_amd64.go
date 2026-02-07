//go:build windows && amd64

package ffmpeg

import (
	"context"
	_ "embed"
)

//go:embed ffmpeg_binaries/windows_amd64/ffmpeg.exe
var ffmpegBinary []byte

//go:embed ffmpeg_binaries/windows_amd64/ffprobe.exe
var ffprobeBinary []byte

func runFFmpeg(ctx context.Context, args ...string) (string, string, error) {
	return runFFmpegInternal(ctx, ffmpegBinary, "ffmpeg.exe", args...)
}

func runFFprobe(ctx context.Context, args ...string) (string, string, error) {
	return runFFmpegInternal(ctx, ffprobeBinary, "ffprobe.exe", args...)
}

func runFFmpegProgress(ctx context.Context, onProgress func(string), args ...string) (string, string, error) {
	return runFFmpegProgressInternal(ctx, ffmpegBinary, "ffmpeg.exe", onProgress, args...)
}
