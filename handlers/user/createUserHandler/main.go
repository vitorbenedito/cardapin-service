package main

import (
	"encoding/json"

	"cardap.in/lambda/auth"
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
	service := services.UserService{}
	userRequest := &model.UserRequestJSON{}
	if loggedUser, err := auth.TokenValid(req.Headers[httphelper.AuthorizationHeader]); err != nil || !loggedUser.IsAdmin() {
		resp.StatusCode = err.Code
		resp.Body = err.Message
		return
	}
	err := json.Unmarshal([]byte(jsonRequest), userRequest)
	if err != nil {
		return httphelper.HandleLambdaResponse(nil, resp, err)
	}
	userJSON, saveErr := service.SaveUser(userRequest.AsModel())
	if saveErr != nil {
		return httphelper.HandleLambdaResponse(nil, resp, saveErr)
	}

	return httphelper.HandleLambdaResponse(userJSON, resp, nil)
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
