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
	var userJSON model.UserLoginJSON
	jsonRequest := req.Body
	err := json.Unmarshal([]byte(jsonRequest), &userJSON)
	userService := services.UserService{}
	userResponseJSON, err := userService.Login(userJSON.AsModel())
	if err != nil {
		resp.StatusCode = 401
		return
	}
	token, _ := auth.CreateToken(userResponseJSON)
	json, err := json.Marshal(userResponseJSON)
	resp.Body = string(json)
	resp.Headers[httphelper.AuthorizationHeader] = "Bearer " + token
	resp.Headers[httphelper.ContentTypeHeader] = httphelper.ApplicationJSONValue
	resp.Headers[httphelper.AccessControlExposeHeaders] = httphelper.AuthorizationHeader
	resp.StatusCode = 200
	return resp, nil
}

func main() {
	handlers.InitCtx()
	lambda.Start(HandleRequest)
}
