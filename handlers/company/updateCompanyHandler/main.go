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
	service := services.CompanyServices{}
	companyJson := &model.CompanyJson{}
	err := json.Unmarshal([]byte(jsonRequest), companyJson)
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}
	company, saveErr := service.Update(*companyJson.AsModel())
	if saveErr != nil {
		return httphelper.HandleLambdaResponse(nil, resp, saveErr)
	}

	return httphelper.HandleLambdaResponse(company.AsJson(), resp, nil)
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
