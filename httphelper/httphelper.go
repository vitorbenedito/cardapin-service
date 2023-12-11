package httphelper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"cardap.in/lambda/apperrors"
	"cardap.in/lambda/model"

	"github.com/aws/aws-lambda-go/events"
)

const (
	AuthorizationHeader                = "x-cardapin-auth-token"
	ContentTypeHeader                  = "Content-Type"
	ApplicationJSONValue               = "application/json"
	AccessControlExposeHeaders         = "Access-Control-Expose-Headers"
	AccessControlAllowOrigin           = "Access-Control-Allow-Origin"
	AccessControlAllowOriginValue      = "http://cardap.in.s3-website.us-east-2.amazonaws.com, https://cardap.in, https://admin.cardap.in, http://cardap.in"
	AccessControlAllowCredentials      = "Access-Control-Allow-Credentials"
	AccessControlAllowCredentialsValue = "true"
)

func EnableCors(req events.APIGatewayProxyRequest, resp *events.APIGatewayProxyResponse) {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.Headers[AccessControlAllowOrigin] = req.Headers["origin"]
	resp.Headers[AccessControlAllowCredentials] = AccessControlAllowCredentialsValue
}

func HandleLambdaResponseEmptySlice(modelSize int, model interface{}, resp events.APIGatewayProxyResponse, err error) (events.APIGatewayProxyResponse, error) {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.Headers[ContentTypeHeader] = ApplicationJSONValue

	if err == nil && modelSize == 0 {
		resp.Body = "[]"
		resp.StatusCode = 404
		return resp, nil
	}
	if err != nil {
		resp.Body = `{"message":"` + err.Error() + `","status": 500}`
		resp.StatusCode = 500
		return resp, nil
	}

	json, errToMarshal := json.Marshal(model)

	if errToMarshal != nil {
		resp.Body = `{"message":"` + errToMarshal.Error() + `","status": 500}`
		resp.StatusCode = 500
		return resp, nil
	}
	return HandleLambdaResponseJson(string(json), resp)
}

func HandleLambdaResponse(model model.RootJSON, resp events.APIGatewayProxyResponse, err error) (events.APIGatewayProxyResponse, error) {
	resp.Headers[ContentTypeHeader] = ApplicationJSONValue

	if err != nil {
		resp.Body = `{"message":"` + err.Error() + `","status": 500}`
		resp.StatusCode = 500
		return resp, nil
	}

	json, errToMarshal := json.Marshal(model)

	if errToMarshal != nil {
		resp.Body = `{"error":"` + errToMarshal.Error() + `","status": 500}`
		resp.StatusCode = 500
		return resp, nil
	}

	if model.GetId() == 0 {
		resp.Body = "{}"
		resp.StatusCode = 404
		return resp, nil
	}

	return HandleLambdaResponseJson(string(json), resp)
}

func HandleLambdaResp(model model.RootJSON, resp events.APIGatewayProxyResponse, appError *apperrors.AppError) (events.APIGatewayProxyResponse, error) {
	resp.Headers[ContentTypeHeader] = ApplicationJSONValue

	if appError != nil {
		resp.Body = `{"message":"` + appError.Error.Error() + `","status":` + fmt.Sprint(appError.Code) + "}"
		resp.StatusCode = appError.Code
		return resp, nil
	}

	json, errToMarshal := json.Marshal(model)

	if errToMarshal != nil {
		resp.Body = `{"error":"` + errToMarshal.Error() + `","status": 500}`
		resp.StatusCode = 500
		return resp, nil
	}

	if model.GetId() == 0 {
		resp.Body = "{}"
		resp.StatusCode = 404
		return resp, nil
	}

	return HandleLambdaResponseJson(string(json), resp)
}

func HandleLambdaResponseJson(json string, resp events.APIGatewayProxyResponse) (events.APIGatewayProxyResponse, error) {

	resp.Headers[ContentTypeHeader] = ApplicationJSONValue

	resp.Body = json
	resp.StatusCode = 200

	return resp, nil
}

func HasConflictLambda(model model.ConflictChecker, resp *events.APIGatewayProxyResponse) bool {
	if err, params := model.HasConflict(); err != nil {
		resp.StatusCode = http.StatusConflict
		resp.Body = `{"message":"` + err.Error() + `","params":["` + strings.Join(params, "\",\"") + `"]}`
		resp.Headers[ContentTypeHeader] = ApplicationJSONValue
		log.Printf("Has conflict: " + resp.Body)
		return true
	}
	return false
}

func HandleResponse(model model.RootJSON, resp http.ResponseWriter, err error) {
	resp.Header().Add(ContentTypeHeader, ApplicationJSONValue)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("{ \"message\": \"" + err.Error() + "\"}"))
		return
	}
	if model.GetId() == 0 {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("{}"))
		return
	}
	json.NewEncoder(resp).Encode(model)
}

func HasConflict(model model.ConflictChecker, resp http.ResponseWriter) bool {
	if err, params := model.HasConflict(); err != nil {
		resp.Header().Add(ContentTypeHeader, ApplicationJSONValue)
		resp.WriteHeader(http.StatusConflict)
		resp.Write([]byte(`{"message":"` + err.Error() + `","params":["` + strings.Join(params, "\",\"") + `"]}`))
		return true
	}
	return false
}

func HandleEmptySliceResponse(modelSize int, resp http.ResponseWriter, writeJson func(resp http.ResponseWriter, jsonModel interface{}), jsonModel interface{}, err error) {
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("{ \"message:\"" + err.Error() + "}"))
		return
	}
	if modelSize == 0 {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("[]"))
		return
	}
	writeJson(resp, jsonModel)
}

func HandleEmptySliceResp(modelSize int, resp http.ResponseWriter, writeJson func(resp http.ResponseWriter, jsonModel interface{}), jsonModel interface{}, appError *apperrors.AppError) {
	if appError != nil {
		resp.WriteHeader(appError.Code)
		resp.Write([]byte("{ \"message:\"" + appError.Error.Error() + "}"))
		return
	}
	if modelSize == 0 {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("[]"))
		return
	}
	writeJson(resp, jsonModel)
}

func GetToken(req *http.Request) string {
	return req.Header.Get(AuthorizationHeader)
}

func GetTokenLambda(req events.APIGatewayProxyRequest) string {
	return req.Headers[AuthorizationHeader]
}
