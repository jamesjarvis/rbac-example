#!/bin/bash

curl -H "User: alice" -X POST http://localhost:8080/v1/map/test -d "pass"
curl -H "User: alice" -X GET http://localhost:8080/v1/map/test
curl -H "User: alice" -X POST http://localhost:8080/v1/map/test2 -d "pass"
curl -H "User: charli" -X GET http://localhost:8080/v1/map/test2

HTTP_STATUS=$(curl -o /dev/null -s -w "%{http_code}\n" -H "User: bob" -X GET http://localhost:8080/v1/map/test2)
if [ "$HTTP_STATUS" -eq "401" ]; then
    echo "pass"
else
    echo "$HTTP_STATUS"
fi

HTTP_STATUS=$(curl -o /dev/null -s -w "%{http_code}\n" -H "User: charli" -X POST http://localhost:8080/v1/map/test3 -d "fail")
if [ "$HTTP_STATUS" -eq "401" ]; then
    echo "pass"
else
    echo "$HTTP_STATUS"
fi