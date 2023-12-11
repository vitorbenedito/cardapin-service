package main

import (
	"log"

	"encoding/json"

	"cardap.in/lambda/email"
	"cardap.in/lambda/httphelper"

	"cardap.in/lambda/services"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, e error) {
	httphelper.EnableCors(req, &resp)
	jsonRequest := req.Body
	service := services.CompanyServices{}
	email := email.Email{}
	err := json.Unmarshal([]byte(jsonRequest), &email)
	if err != nil {
		resp.Body = `{"message":"` + err.Error() + `","status": 500}`
		resp.StatusCode = 500
		return resp, nil
	}

	log.Printf("Sending mail to: " + email.Email)
	if sent := service.CompanyInterested(email); !sent {
		resp.StatusCode = 500
		return resp, nil
	}
	resp.StatusCode = 200
	return resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
