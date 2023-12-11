package main

import (
	"strconv"

	"cardap.in/lambda/httphelper"

	"cardap.in/lambda/services"

	"cardap.in/lambda/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, e error) {
	httphelper.EnableCors(req, &resp)
	id := req.PathParameters["phone"]
	service := services.ClientService{}
	phone, _ := strconv.ParseUint(id, 10, 64)
	json, err := service.GetByPhone(uint64(phone))
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}

	return httphelper.HandleLambdaResponse(json, resp, nil)
}
func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
