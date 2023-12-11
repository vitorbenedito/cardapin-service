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
	service := services.TableServices{}
	table := model.TableJSON{}
	err := json.Unmarshal([]byte(jsonRequest), &table)

	tableToSave := table.AsModel()
	if httphelper.HasConflictLambda(tableToSave, &resp) {
		return resp, nil
	}
	savedTable, err := service.Save(*tableToSave, false)
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}
	return httphelper.HandleLambdaResponse(savedTable.AsJSON(), resp, nil)
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
