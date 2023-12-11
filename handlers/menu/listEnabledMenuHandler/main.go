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
	id := req.PathParameters["companyCode"]
	service := services.MenuServices{}
	menuJSON, saveErr := service.GetMenuEnabledByCompanyCode(id)
	if saveErr != nil {
		return httphelper.HandleLambdaResponse(nil, resp, saveErr)
	}

	return httphelper.HandleLambdaResponse(menuJSON, resp, nil)
}
func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
