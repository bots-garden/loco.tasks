#!/bin/bash

echo "💥 Stopping all processes of ${TASK_GROUP}"
curl -X DELETE ${LOCO_TASKS_PROCESSES}/${TASK_GROUP}
