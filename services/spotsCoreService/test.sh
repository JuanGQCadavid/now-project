#!/bin/bash

TEMP_DIR=.tempWorkDir
APP_NAME=spotscore

mkdir ${TEMP_DIR}
cp -R ../pkgs ${TEMP_DIR}

docker build -t ${APP_NAME} .

rm -rf ${TEMP_DIR}

docker run -it -p 8000:8000 -v $HOME/.aws:/root/.aws --env-file .env ${APP_NAME}