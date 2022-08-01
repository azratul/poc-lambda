package main

import (
	"context"
	"lambda-parent/config"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

// init is automatically run upon Lambda cold start
func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	if os.Getenv("LAMBDA_POSTS") == "" || os.Getenv("LAMBDA_USERS") == "" || os.Getenv("REGION") == "" {
		log.Println("All env vars must be set!!!")
		return
	}

	config.Conf.AWS.LambdaPosts = os.Getenv("LAMBDA_POSTS")
	config.Conf.AWS.LambdaUsers = os.Getenv("LAMBDA_USERS")
	config.Conf.AWS.Region = os.Getenv("REGION")

	r := CreateServer()
	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Request received:\n%+v", req)
	return ginLambda.ProxyWithContext(ctx, req)
}
