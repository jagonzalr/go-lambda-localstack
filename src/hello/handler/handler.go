package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

var okResponse = Response{
	StatusCode:      200,
	IsBase64Encoded: false,
	Headers: map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
		"Cache-Control":                    "no-cache; no-store",
		"Content-Type":                     "application/json",
		"Content-Security-Policy":          "default-src self",
		"Strict-Transport-Security":        "max-age=31536000; includeSubDomains",
		"X-Content-Type-Options":           "nosniff",
		"X-XSS-Protection":                 "1; mode=block",
		"X-Frame-Options":                  "DENY",
	},
}

type httpHandler struct {
}

// Handler - interface
type Handler interface {
	Run(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error)
}

func (h httpHandler) Run(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error) {
	okResponse.Body = "Hello World"
	return okResponse, nil
}

// NewHttpHandler -
func NewHttpHandler() *httpHandler {
	return &httpHandler{}
}
