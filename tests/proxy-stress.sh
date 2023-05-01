#!/bin/bash
# -n 10000 -c 1000
#hey -n 10 -c 5 -m POST \
hey -n 300 -c 100 -m POST \
-H "Content-Type: application/json; charset=utf-8" \
-d '{"name":"Bob Morane","age":42}' \
"http://localhost:8080/functions/hello-world"
