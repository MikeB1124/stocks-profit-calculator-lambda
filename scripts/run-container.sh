#!/bin/bash

docker stop stocks-profit-calculator-lambda
docker rm stocks-profit-calculator-lambda
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go
docker build -t stocks-lambda-image .
docker run --name stocks-profit-calculator-lambda -p 9000:8080 --env-file .env stocks-lambda-image