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
	service := services.MenuServices{}
	if ok, err := service.DeleteMenu(id); !ok {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}
	resp.StatusCode = 204
	return resp, nil
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
