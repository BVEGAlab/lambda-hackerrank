package lambdas

import (
    "context"


    "github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (string, error) {
    // in development
    return "hello world", nil
}
