package httphelper

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

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

func HandleResponse(model model.RootJSON, c *gin.Context, err error) {
	c.Header(ContentTypeHeader, ApplicationJSONValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "{ \"message\": \""+err.Error()+"\"}")
		return
	}
	if model.GetId() == 0 {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.Writer.Write([]byte("{}"))
		return
	}
	c.JSON(http.StatusOK, model)
}

func HasConflict(model model.ConflictChecker, c *gin.Context) bool {
	if err, params := model.HasConflict(); err != nil {
		c.Header(ContentTypeHeader, ApplicationJSONValue)
		c.JSON(http.StatusConflict, `{"message":"`+err.Error()+`","params":["`+strings.Join(params, "\",\"")+`"]}`)
		return true
	}
	return false
}

func HandleEmptySliceResponse(modelSize int, c *gin.Context, jsonModel interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, "{ \"message\": \""+err.Error()+"\"}")
		return
	}
	if modelSize == 0 {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.Writer.Write([]byte("[]"))
		return
	}
	c.JSON(http.StatusOK, jsonModel)
}

func HandleEmptySliceResp(modelSize int, c *gin.Context, jsonModel interface{}, appError *apperrors.AppError) {
	if appError != nil {
		c.JSON(http.StatusInternalServerError, "{ \"message\": \""+appError.Error.Error()+"\"}")
		return
	}
	if modelSize == 0 {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.Writer.Write([]byte("[]"))
		return
	}
	c.JSON(http.StatusOK, jsonModel)
}

func GetToken(req *http.Request) string {
	return req.Header.Get(AuthorizationHeader)
}
