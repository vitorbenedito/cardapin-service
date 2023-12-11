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
	service := services.ClientService{}
	client := &model.Client{}
	err := json.Unmarshal([]byte(jsonRequest), client)
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}
	_, saveErr := service.Save(*client, true)
	if saveErr != nil {
		return httphelper.HandleLambdaResponse(nil, resp, saveErr)
	}

	return httphelper.HandleLambdaResponse(client, resp, nil)
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
