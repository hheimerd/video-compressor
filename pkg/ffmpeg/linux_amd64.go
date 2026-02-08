//go:build linux && amd64

package ffmpeg

import (
	"context"
	_ "embed"
)

//go:embed ffmpeg_binaries/linux_amd64/ffmpeg
var ffmpegBinary []byte

func runFFmpeg(ctx context.Context, args ...string) (string, string, error) {
	return runFFmpegInternal(ctx, ffmpegBinary, "ffmpeg", args...)
}

func runFFmpegProgress(ctx context.Context, onProgress func(string), args ...string) (string, string, error) {
	return runFFmpegProgressInternal(ctx, ffmpegBinary, "ffmpeg", onProgress, args...)
}
