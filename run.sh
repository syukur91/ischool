#!/bin/bash

echo "building container-mgmt ..."

if CGO_ENABLED=0 go build -o ./bin/ischool
    then
        echo "completed building ischool"

      ./bin/ischool
    else
        echo "failed to build ischool"
fi