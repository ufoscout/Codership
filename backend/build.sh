#!/bin/bash

PROJECT_NAME=backend
BUILD_TARGET_PATH=.
BUILD_TARGET_FILE=backend

declare -a steps=(
  #"rm -rf ${BUILD_TARGET_PATH}"
  "GOCACHE=off vgo test ./... -v"
  "vgo build -o ${BUILD_TARGET_PATH}/${BUILD_TARGET_FILE}"
)

echo $LINE_SEPARATOR
echo "Build $PROJECT_NAME"
echo $LINE_SEPARATOR

for i in "${steps[@]}"
do
    echo "Execute step: '$i'"
    eval $i
    rc=$?
    if [[ $rc -ne 0 ]] ; then
        echo "Failure executing: $i"; exit $rc
    fi
done
