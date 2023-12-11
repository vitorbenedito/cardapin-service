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
	id := req.PathParameters["id"]
	groupsJSON := model.ListAdditionalItemsByCompanyId(id)
	json, err := json.Marshal(groupsJSON)
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}
	return httphelper.HandleLambdaResponseJson(string(json), resp)
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
