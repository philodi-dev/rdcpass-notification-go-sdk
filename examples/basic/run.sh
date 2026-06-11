#!/usr/bin/env bash
# Work around dyld "missing LC_UUID" on macOS 26 with Go < 1.24.
set -euo pipefail
cd "$(dirname "$0")"
export CGO_ENABLED=0
exec go run .
