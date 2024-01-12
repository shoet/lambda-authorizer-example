package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/shoet/lambda-authorizer-example/internal/interfaces"
)

var echoLambdaHTTP *echoadapter.EchoLambdaV2

func ExitOnErr(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}

func init() {
	srv, err := interfaces.NewServer()
	if err != nil {
		ExitOnErr(err)
	}
	echoLambdaHTTP = echoadapter.NewV2(srv)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	fmt.Println("come handler")
	h := req.Headers
	fmt.Println(h)
	return echoLambdaHTTP.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
