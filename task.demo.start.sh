#!/bin/bash
curl -X POST http://localhost:7070/admin/tasks/registration \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"start-hello",
    "description":"start-hello service",
    "path":"./little.process.sh",
    "args": ["main-programm", "👋 hello", "world 🌍"],
    "env": ["FILE_NAME=tada.txt"]
}
EOF
