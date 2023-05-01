#!/bin/bash

<<DOCUMENTATION
    -------------------------------
    Stop a group of task.
    -------------------------------
DOCUMENTATION

TASK_GROUP="hello-world-group"

FUNCTION_NAME="hello-world"
FUNCTION_REVISION="default"

LOCO_TASKS_START="http://localhost:7070/admin/tasks/start"
LOCO_TASKS_PROCESSES="http://localhost:7070/admin/tasks/processes"
LOCO_SWITCH_REGISTRATION="http://localhost:8080/admin/functions/registration"

curl -X POST ${LOCO_TASKS_START} \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "group":"${TASK_GROUP}",
    "description":"stopping all processes of ${TASK_GROUP}",
    "path":"./services/capsule/tasks.kill.sh",
    "args": [],
    "env": [
        "LOCO_TASKS_PROCESSES=${LOCO_TASKS_PROCESSES}",
        "TASK_GROUP=${TASK_GROUP}"
    ]
}
EOF
echo ""

# unregister the function from the proxy

curl -X DELETE ${LOCO_SWITCH_REGISTRATION} \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"${FUNCTION_NAME}",
    "revision":"${FUNCTION_REVISION}"
}
EOF
echo ""
