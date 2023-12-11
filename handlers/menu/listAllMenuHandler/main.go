package main

import (
	"encoding/json"

	"cardap.in/lambda/httphelper"

	"cardap.in/lambda/services"

	"cardap.in/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, e error) {
	httphelper.EnableCors(req, &resp)
	service := services.MenuServices{}
	menuJSON, saveErr := service.GetMenuByLoggedCompany(httphelper.GetTokenLambda(req))
	if saveErr != nil {
		return httphelper.HandleLambdaResp(nil, resp, saveErr)
	}
	json, err := json.Marshal(menuJSON)
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}
	return httphelper.HandleLambdaResponseJson(string(json), resp)
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
