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
	service := services.ImageServices{}
	fileJSON := model.FileRequest{}
	err := json.Unmarshal([]byte(jsonRequest), &fileJSON)
	json, _ := service.GeneratePresignedUrlToPut(fileJSON)
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}

	return httphelper.HandleLambdaResponseJson(json, resp)
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
