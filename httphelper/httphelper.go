package httphelper

import (
	"encoding/json"
	"net/http"
	"strings"

	"cardap.in/apperrors"
	"cardap.in/model"
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
