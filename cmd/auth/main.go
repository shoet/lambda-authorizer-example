package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const TOKEN_KEY = "x-api-token"

func Handler(
	ctx context.Context, req events.APIGatewayCustomAuthorizerRequestTypeRequest,
) (events.APIGatewayCustomAuthorizerResponse, error) {
	fmt.Println("come auth")
	fmt.Println(req.Headers)
	token, ok := req.Headers[TOKEN_KEY]
	if !ok {
		return events.APIGatewayCustomAuthorizerResponse{}, fmt.Errorf("token not found")
	}
	if token != "XXXXXDUMMYTOKENXXXXX" {
		return events.APIGatewayCustomAuthorizerResponse{}, fmt.Errorf("token not match")
	}
	fmt.Println("auth success")
	return events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: "user",
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   "Allow",
					Resource: []string{"arn:aws:execute-api:*:*:*"},
				},
			},
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
