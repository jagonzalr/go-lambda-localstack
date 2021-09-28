package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jagonzalr/go-lambda-localstack/src/hello/handler"
)

func main() {
	handler := handler.Create()
	lambda.Start(handler.Run)
}
