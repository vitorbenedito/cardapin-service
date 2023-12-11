package main

import (
	"cardap.in/lambda/httphelper"
	"cardap.in/lambda/services"

	"cardap.in/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, e error) {
	httphelper.EnableCors(req, &resp)
	id := req.PathParameters["id"]
	service := services.UserService{}
	user, err := service.GetUserById(id)
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}

	return httphelper.HandleLambdaResponse(user, resp, nil)
}
func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
