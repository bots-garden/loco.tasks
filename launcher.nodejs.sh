#!/bin/bash
FUNCTION_NAME=${FUNCTION_NAME}

# Register the function to the proxy
echo "🤖 register the function to the proxy"
LOCO_SWITCH_REGISTRATION="http://localhost:8080/admin/functions/registration"

curl -X POST ${LOCO_SWITCH_REGISTRATION} \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"${FUNCTION_NAME}",
    "revision":"default",
    "httpPort":${HTTP_PORT},
    "status":0,
    "https":false,
    "domain":"localhost"
}
EOF

# Download wasm file
# TODO

# Start the WASM server
echo "🚀 start the server"
HTTP_PORT="${HTTP_PORT}" \
WASM_FILE="${WASM_FILE}" \
node --experimental-wasi-unstable-preview1 --no-warnings ../loco.services/hello/index.mjs

