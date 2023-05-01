#!/bin/bash

source ./ephemeral.port.sh
HTTP_PORT=$(ephemeral_port)
WASM_FILE="./services/capsule/hello-world.wasm"

FUNCTION_NAME="hello-world"
FUNCTION_REVISION="default"
FUNCTION_DOMAIN="localhost"

TASK_GROUP="hello-world-group"

LOCO_SWITCH_REGISTRATION="http://localhost:8080/admin/functions/registration"
LOCO_TASKS_START="http://localhost:7070/admin/tasks/start"


# Start the function with the scheduler
curl -X POST ${LOCO_TASKS_START} \-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "group":"${TASK_GROUP}",
    "name":"function ${FUNCTION_NAME} ${FUNCTION_REVISION} ${HTTP_PORT}",
    "description":"${FUNCTION_NAME} ${FUNCTION_REVISION} ${HTTP_PORT}",
    "path":"./services/capsule/capsule-http",
    "args": ["", "-wasm=${WASM_FILE}", "-httpPort=${HTTP_PORT}"],
    "env": []
}
EOF


# Register the function to the proxy
curl -X POST ${LOCO_SWITCH_REGISTRATION} \
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

# Start the WASM server
echo "🚀 start ${FUNCTION_NAME} [${FUNCTION_REVISION}] on port ${HTTP_PORT}"


<<DOCUMENTATION

DOCUMENTATION


