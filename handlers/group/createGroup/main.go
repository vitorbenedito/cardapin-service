package main

import (
	"encoding/json"

	"cardap.in/lambda/httphelper"

	"cardap.in/lambda/model"

	"cardap.in/lambda/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, e error) {
	httphelper.EnableCors(req, &resp)
	jsonRequest := req.Body
	additionalItemGroupJSON := &model.AdditionalItemsGroupJSON{}
	err := json.Unmarshal([]byte(jsonRequest), &additionalItemGroupJSON)

	additionalItemGroup := additionalItemGroupJSON.AsModel()
	if httphelper.HasConflictLambda(additionalItemGroup, &resp) {
		return resp, nil
	}

	savedGroup, err := model.SaveAdditionalItemsGroup(*additionalItemGroup)
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}
	return httphelper.HandleLambdaResponse(savedGroup.AsJSON(), resp, nil)
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
