#!/bin/bash

# Define versions and base URL
FFMPEG_VERSION="6.1"
BASE_URL="https://github.com/ffbinaries/ffbinaries-prebuilt/releases/download/v${FFMPEG_VERSION}"

# Define target directory
TARGET_DIR="$(dirname "$0")/../pkg/ffmpeg/ffmpeg_binaries"

# Create target directory if it doesn't exist
mkdir -p "$TARGET_DIR"

# Function to download and extract
download_and_extract() {
    local platform="$1"
    local target_platform_dir="$2"
    local ext="$3"
    
    echo "Processing $platform..."
    
    # Create platform specific directory
    mkdir -p "$TARGET_DIR/$target_platform_dir"
    
    # Download ffmpeg
    echo "Downloading ffmpeg for $platform..."
    curl -L "$BASE_URL/ffmpeg-$FFMPEG_VERSION-$platform.zip" -o "ffmpeg.zip"
    unzip -o -j "ffmpeg.zip" -d "$TARGET_DIR/$target_platform_dir"
    rm "ffmpeg.zip"
    
    # Download ffprobe
    echo "Downloading ffprobe for $platform..."
    curl -L "$BASE_URL/ffprobe-$FFMPEG_VERSION-$platform.zip" -o "ffprobe.zip"
    unzip -o -j "ffprobe.zip" -d "$TARGET_DIR/$target_platform_dir"
    rm "ffprobe.zip"

    # Rename if necessary (remove extension for non-windows if it exists, though unzip usually handles this)
    if [ "$ext" == ".exe" ]; then
         # Windows files should have .exe, ensure they do (ffbinaries usually have it)
         :
    else
        # Ensure execute permissions for unix-like
        chmod +x "$TARGET_DIR/$target_platform_dir/ffmpeg"
        chmod +x "$TARGET_DIR/$target_platform_dir/ffprobe"
    fi
}

# Download for Darwin AMD64
download_and_extract "macos-64" "darwin_amd64" ""

# Download for Darwin ARM64
# ffbinaries for 6.1 uses "macos-64". It seems they might not have a separate arm64 build published,
# or macos-64 covers it (likely x64 running via Rosetta).
echo "Copying macos-64 to darwin_arm64..."
mkdir -p "$TARGET_DIR/darwin_arm64"
cp "$TARGET_DIR/darwin_amd64/ffmpeg" "$TARGET_DIR/darwin_arm64/ffmpeg"
cp "$TARGET_DIR/darwin_amd64/ffprobe" "$TARGET_DIR/darwin_arm64/ffprobe"


# Download for Linux AMD64
download_and_extract "linux-64" "linux_amd64" ""

# Download for Windows AMD64
download_and_extract "win-64" "windows_amd64" ".exe"

echo "Done! Binaries are in $TARGET_DIR"
