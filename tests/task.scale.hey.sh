#!/bin/bash

source ./ephemeral.port.sh
HTTP_PORT=$(ephemeral_port)
WASM_FILE="./services/capsule/hey.wasm"

FUNCTION_NAME="hey"
FUNCTION_REVISION="default"
FUNCTION_DOMAIN="localhost"

TASK_GROUP="hey-group"

LOCO_SWITCH_ENDPOINT="http://localhost:8080/admin/functions/endpoint"
LOCO_TASKS_START="http://localhost:7070/admin/tasks/start"


# Start a new instance of the function with the scheduler
curl -X POST ${LOCO_TASKS_START} \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "group":"${TASK_GROUP}",
    "name":"function ${FUNCTION_NAME} ${FUNCTION_REVISION} ${HTTP_PORT}",
    "description":"[scaled] ${FUNCTION_NAME} ${FUNCTION_REVISION} ${HTTP_PORT}",
    "path":"./services/capsule/capsule-http",
    "args": ["", "-wasm=${WASM_FILE}", "-httpPort=${HTTP_PORT}"],
    "env": []
}
EOF

# 🚀 scale the function to the proxy"
# Register the function to the proxy

curl -X POST ${LOCO_SWITCH_ENDPOINT} \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"${FUNCTION_NAME}",
    "revision":"${FUNCTION_REVISION}",
    "httpPort":${HTTP_PORT},
    "status":0,
    "https":false,
    "domain":"${FUNCTION_DOMAIN}"
}
EOF


<<DOCUMENTATION

DOCUMENTATION
