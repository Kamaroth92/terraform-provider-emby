#!/usr/bin/env bash
# Copyright (c) HashiCorp, Inc.

# Usage:
#   ./generate-sdk.sh -h "https://my-emby-server:8096"

set -euo pipefail

usage() {
    echo "Usage: $(basename "$0") -h <emby_host>"
    echo "  -h  Emby server URL (e.g. https://my-emby-server:8096)"
    exit 1
}

while getopts "h:" opt; do
    case "$opt" in
        h) EMBY_HOST="$OPTARG" ;;
        *) usage ;;
    esac
done

if [ -z "${EMBY_HOST:-}" ]; then
    echo "Error: -h <emby_host> is required"
    usage
fi

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
SPEC_URL="$EMBY_HOST/emby/openapi"
SPEC_FILE="$SCRIPT_DIR/openapi-spec.json"
OUTPUT_DIR="$SCRIPT_DIR/internal/client"

echo "Downloading OpenAPI spec from $SPEC_URL ..."
curl -sS -o "$SPEC_FILE" "$SPEC_URL"
echo "Saved spec to $SPEC_FILE"

echo "Building codegen Docker image..."
docker build -f "$SCRIPT_DIR/Dockerfile.codegen" -t emby-go-codegen "$SCRIPT_DIR"

if [ -d "$OUTPUT_DIR" ]; then
    echo "Clearing old generated files in $OUTPUT_DIR ..."
    rm -rf "$OUTPUT_DIR"/*
else
    mkdir -p "$OUTPUT_DIR"
fi

echo "Generating Go SDK..."
docker run --rm \
    -v "${SPEC_FILE}:/spec/openapi.json" \
    -v "${OUTPUT_DIR}:/out" \
    emby-go-codegen \
    -i /spec/openapi.json

echo "Fixing ServerConfiguration type collision in configuration.go ..."
CONFIG_FILE="$OUTPUT_DIR/configuration.go"
sed -i \
    -e 's/type ServerConfiguration struct/type OAPIServerConfig struct/' \
    -e 's/type ServerConfigurations \[\]ServerConfiguration/type OAPIServerConfigs []OAPIServerConfig/' \
    -e 's/ServerConfigurations/OAPIServerConfigs/g' \
    "$CONFIG_FILE"

echo ""
echo "Done! Generated SDK is in $OUTPUT_DIR"
