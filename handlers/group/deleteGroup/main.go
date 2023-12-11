package main

import (
	"cardap.in/lambda/httphelper"
	"cardap.in/lambda/model"

	"cardap.in/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, e error) {
	httphelper.EnableCors(req, &resp)
	id := req.PathParameters["id"]
	err := model.DeleteAdditionalGroup(id)
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}
	resp.StatusCode = 204
	return resp, nil
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
