#!/bin/bash -e

APP_ENV=${APP_ENV:-local}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment..."

CONFIG_FILE=./config/${APP_ENV}.yaml

if [[ -z ${APP_DSN} ]]; then
  export APP_DSN=`sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' ${CONFIG_FILE}`
fi

echo "[`date`] Running DB migrations..."
migrate -database "${APP_DSN}" -path ./migrations up

echo "[`date`] Starting server..."
./service -p ${CONFIG_FILE} 
