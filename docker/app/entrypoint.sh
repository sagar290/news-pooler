#!/bin/sh
set -e

if [ ${APP_ENV} = 'local' ];
then
    echo "not in production"
    cd /src && air
else
    echo "in production"
     (cd /src/) && (go build main.go) && (/src/./main)
fi