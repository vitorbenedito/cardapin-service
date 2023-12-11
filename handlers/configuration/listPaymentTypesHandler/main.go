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
	service := services.PaymentTypeService{}
	paymentTypes, saveErr := service.ListPaymentTypes()

	if saveErr != nil {
		return httphelper.HandleLambdaResponse(nil, resp, saveErr)
	}

	return httphelper.HandleLambdaResponseEmptySlice(len(paymentTypes), paymentTypes, resp, nil)

}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
