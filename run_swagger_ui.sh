#!/bin/sh

docker run -p 80:8080 -e SWAGGER_JSON=/foo/swagger.json -v /home/gabriel/go_project/src:/foo swaggerapi/swagger-ui