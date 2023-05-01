# loco.tasks


HTTP_PORT="5555" \
WASM_FILE="../loco.services/hello/main.wasm" \
node --experimental-wasi-unstable-preview1 --no-warnings ../loco.services/hello/index.mjs

curl -X POST http://localhost:5555 \
-H 'Content-Type: text/plain; charset=utf-8' \
-d 'Jane Doe'
echo ""

curl -X POST http://localhost:8080/functions/hello/default \
-H 'Content-Type: text/plain; charset=utf-8' \
-d 'Bob Morane'
echo ""

launcher.sh
