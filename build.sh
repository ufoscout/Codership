#!/bin/bash

## Backend
cd backend
./build.sh
rc=$?
if [[ $rc -ne 0 ]] ; then
  echo 'Failure'; cd ..; exit $rc
fi
cd ..

## Frontend
cd frontend
./build.sh
rc=$?
if [[ $rc -ne 0 ]] ; then
  echo 'Failure'; cd ..; exit $rc
fi
cd ..
