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

var echoLambda *echoadapter.EchoLambda

func ExitOnErr(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}

func init() {
	srv, err := interfaces.NewServer()
	if err != nil {
		ExitOnErr(err)
	}
	echoLambda = echoadapter.New(srv)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
