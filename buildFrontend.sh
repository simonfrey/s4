#!/bin/bash


echo "[1] Build wasm"
cd wasm || exit
GOOS=js GOARCH=wasm go build -o wasm.wasm wasm.go
cd .. || exit

echo "[2] Pack wasm into javascript"

echo 'var wasmPayloadInlineURL = `data:application/wasm;base64,' > frontend/js/wasmPayload.js
base64 wasm/wasm.wasm >> frontend/js/wasmPayload.js
echo '`' >> frontend/js/wasmPayload.js

echo "[3] Done. You can upload the frontend folder to your server or use it locally"