#!/bin/bash

# Get the absolute path to the script's directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# Get the project root directory (assuming scripts/ is one level deep)
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Navigate to the project root
cd "$PROJECT_ROOT" || exit 1

echo "Installing Go dependencies..."
go mod download
go mod tidy

echo "Downloading external binaries..."
"$SCRIPT_DIR/download_binaries.sh"

echo "Setup completed successfully."
