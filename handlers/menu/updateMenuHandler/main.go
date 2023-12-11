package main

import (
	"encoding/json"

	"cardap.in/lambda/httphelper"

	"cardap.in/lambda/model"
	"cardap.in/lambda/services"

	"cardap.in/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, e error) {
	httphelper.EnableCors(req, &resp)
	jsonRequest := req.Body
	service := services.MenuServices{}
	menuJSON := &model.MenuJSON{}
	err := json.Unmarshal([]byte(jsonRequest), menuJSON)
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}
	menu := *menuJSON.AsModel()
	if httphelper.HasConflictLambda(&menu, &resp) {
		return resp, nil
	}
	menuSavedJSON, saveErr := service.Update(menu)
	if saveErr != nil {
		return httphelper.HandleLambdaResponse(nil, resp, saveErr)
	}

	return httphelper.HandleLambdaResponse(menuSavedJSON, resp, nil)
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
