#!/bin/bash

set -eu

echo "[1] Build wasm"
wasmPayloadFile=$(mktemp)
trap 'rm -f $wasmPayloadFile' ERR EXIT
GOOS=js GOARCH=wasm go build -o "$wasmPayloadFile" wasm/wasm.go

echo "[2] Pack wasm into javascript"
wasmJsPayloadFile="frontend/js/wasmPayload.js"
echo 'const wasmPayloadInlineURL = `data:application/wasm;base64,' > "$wasmJsPayloadFile"
base64 "$wasmPayloadFile" >> "$wasmJsPayloadFile"
echo '`' >> "$wasmJsPayloadFile"

echo "[3] Done. You can upload the frontend folder to your server or use it locally"