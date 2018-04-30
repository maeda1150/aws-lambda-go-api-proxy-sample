package main

import (
	"aws-lambda-go-api-proxy-sample/controllers"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var initialized = false
var ginLambda *ginadapter.GinLambda

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	gin.SetMode(gin.ReleaseMode)

	if !initialized {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")
		r := gin.New()
		v := r.Group("/")
		{
			v.POST("test", controllers.PostSample)
			v.GET("test", controllers.GetSample)
		}

		ginLambda = ginadapter.New(r)
		initialized = true
	}

	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
