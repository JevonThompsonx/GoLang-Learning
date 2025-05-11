#!/bin/bash

# Name of your Go project binary
BINARY_NAME="myapp"

# Output directory
OUTPUT_DIR="build"

# Platforms and architectures to build for
PLATFORMS=("linux" "darwin" "windows")
ARCHITECTURES=("amd64" "arm64")

# Clean output directory
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

# Loop through all OS/ARCH combinations
for GOOS in "${PLATFORMS[@]}"; do
  for GOARCH in "${ARCHITECTURES[@]}"; do
    OUTPUT_NAME="${BINARY_NAME}-${GOOS}-${GOARCH}"
    if [ "$GOOS" == "windows" ]; then
      OUTPUT_NAME+=".exe"
    fi

    echo "Building for $GOOS/$GOARCH..."

    env GOOS=$GOOS GOARCH=$GOARCH go build -o "$OUTPUT_DIR/$OUTPUT_NAME"

    if [ $? -ne 0 ]; then
      echo "❌ Build failed for $GOOS/$GOARCH"
    else
      echo "✅ Build succeeded: $OUTPUT_NAME"
    fi
  done
done

echo "All builds complete. Binaries are in the '$OUTPUT_DIR' directory."
