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
	service := services.SectionService{}
	sections, saveErr := service.ListSection()

	if saveErr != nil {
		return httphelper.HandleLambdaResponse(nil, resp, saveErr)
	}

	return httphelper.HandleLambdaResponseEmptySlice(len(sections), sections, resp, nil)

}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
