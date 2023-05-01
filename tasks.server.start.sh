#!/bin/bash
curl -X POST http://localhost:7070/admin/tasks/registration \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"start-😃",
    "description":"wasm service",
    "path":"./launcher.nodejs.sh",
    "args": [],
    "env": ["HTTP_PORT=5555", "WASM_FILE=../loco.services/hello/main.wasm", "FUNCTION_NAME=hello"]
}
EOF


curl -X POST http://localhost:7070/admin/tasks/registration \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"start-🎃",
    "description":"wasm service",
    "path":"./launcher.nodejs.sh",
    "args": [],
    "env": ["HTTP_PORT=5556", "WASM_FILE=../loco.services/hello/main.wasm", "FUNCTION_NAME=hey"]
}
EOF
