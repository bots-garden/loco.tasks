#!/bin/bash

curl -X POST  http://localhost:8080/functions/hello-world \
-H "Content-Type: application/json; charset=utf-8" \
-d '{"name":"Bob Morane","age":42}'
echo ""

curl -X POST  http://localhost:8080/functions/hello-world \
-H "Content-Type: application/json; charset=utf-8" \
-d '{"name":"Jane Doe","age":24}'
echo ""

curl -X POST  http://localhost:8080/functions/hey \
-H 'Content-Type: text/plain; charset=utf-8' \
-d 'Bob Morane' 
echo ""