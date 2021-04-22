#!/bin/bash

set -eu

echo "[1] Build wasm"
wasmPayloadFile=$(mktemp)
trap 'rm -f $wasmPayloadFile' ERR EXIT
GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o "$wasmPayloadFile" wasm/wasm.go

echo "[2] Pack wasm into javascript"
wasmJsPayloadFile="frontend/js/wasmPayload.js"
echo 'const wasmPayloadInlineURL = `data:application/wasm;base64,' > "$wasmJsPayloadFile"
base64 "$wasmPayloadFile" >> "$wasmJsPayloadFile"
echo '`' >> "$wasmJsPayloadFile"

echo "[3] Pack css & javascript into a single HTML file"
htmlTemplate="frontend/index.template.html"
mkdir -p build
{
  awk '1;/style/{exit}' "$htmlTemplate"
  cat frontend/css/*
  awk '/<\/style/,0;/script/{exit}' "$htmlTemplate"
  cat frontend/js/*
  awk '/<\/script/,0' "$htmlTemplate"
}  > build/index.html

echo "[4] Done. You can upload the build folder to your server or use it locally"