#!/bin/bash

echo "Runing build script"
sh build.sh

echo "Zipping"
zip -r main.zip main

echo "Deploying"
aws lambda update-function-code --function-name LambdaName --zip-file fileb://main.zip

echo "Cleaning"
rm -rf main.zip
rm -rf main