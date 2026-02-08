#!/bin/bash

# Exit on any error
set -e

# clean releases
rm -rf releases
mkdir -p releases

echo "Starting build process for multiple platforms..."

# ------------------------------------------------------------------
# macOS (Intel)
# ------------------------------------------------------------------
echo "Building for darwin/amd64..."
wails build -platform darwin/amd64 -o VideoCompressor-amd64.app

# Copy raw artifact
echo "Copying to releases..."
cd build/bin
if [ -d "VideoCompressor-amd64.app" ]; then
    cp -R VideoCompressor-amd64.app ../../releases/VideoCompressor-darwin-amd64.app
else
    echo "Warning: VideoCompressor-amd64.app not found"
fi
cd ../..

# ------------------------------------------------------------------
# macOS (Apple Silicon)
# ------------------------------------------------------------------
echo "Building for darwin/arm64..."
wails build -platform darwin/arm64 -o VideoCompressor-arm64.app

# Copy raw artifact
echo "Copying to releases..."
cd build/bin
if [ -d "VideoCompressor-arm64.app" ]; then
    cp -R VideoCompressor-arm64.app ../../releases/VideoCompressor-darwin-arm64.app
else
    echo "Warning: VideoCompressor-arm64.app not found"
fi
cd ../..

# ------------------------------------------------------------------
# Linux (AMD64)
# ------------------------------------------------------------------
echo "Building for linux/amd64..."
# Note: Cross-compiling for Linux from macOS requires Docker or meaningful setup.
# Failures here might result in no file being created.
wails build -platform linux/amd64 -o VideoCompressor-linux-amd64 || echo "Linux build command failed or skipped"

# Copy raw artifact
echo "Copying to releases..."
cd build/bin
if [ -f "VideoCompressor-linux-amd64" ]; then
    cp VideoCompressor-linux-amd64 ../../releases/
else
    echo "WARNING: VideoCompressor-linux-amd64 not found in build/bin."
    echo "If you are on macOS, cross-compiling for Linux usually requires Docker."
    echo "See: https://wails.io/docs/guides/cross-compilation"
fi
cd ../..

# ------------------------------------------------------------------
# Windows (AMD64)
# ------------------------------------------------------------------
echo "Building for windows/amd64..."
wails build -platform windows/amd64 -o VideoCompressor-windows-amd64.exe

# Copy raw artifact
echo "Copying to releases..."
cd build/bin
if [ -f "VideoCompressor-windows-amd64.exe" ]; then
    cp VideoCompressor-windows-amd64.exe ../../releases/
else
    echo "Warning: VideoCompressor-windows-amd64.exe not found"
fi
cd ../..

echo "-----------------------------------"
echo "Build procedure finished."
echo "Checking 'releases' folder:"
ls -lh releases/
echo "Checking 'build/bin' contents (for debugging):"
ls -lh build/bin/
