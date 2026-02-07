//go:build darwin && amd64

package ffmpeg

import (
	"context"
	_ "embed"
)

//go:embed ffmpeg_binaries/darwin_amd64/ffmpeg
var ffmpegBinary []byte

//go:embed ffmpeg_binaries/darwin_amd64/ffprobe
var ffprobeBinary []byte

func runFFmpeg(ctx context.Context, args ...string) (string, string, error) {
	return runFFmpegInternal(ctx, ffmpegBinary, "ffmpeg", args...)
}

func runFFprobe(ctx context.Context, args ...string) (string, string, error) {
	return runFFmpegInternal(ctx, ffprobeBinary, "ffprobe", args...)
}

func runFFmpegProgress(ctx context.Context, onProgress func(string), args ...string) (string, string, error) {
	return runFFmpegProgressInternal(ctx, ffmpegBinary, "ffmpeg", onProgress, args...)
}
