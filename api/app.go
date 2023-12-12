package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"cardap.in/auth"
	"cardap.in/db"
	"cardap.in/email"
	"cardap.in/httphelper"
	"cardap.in/migration"
	"cardap.in/model"
	"cardap.in/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type App struct {
	Router *gin.Engine
}

func (a *App) Initialize() {
	a.Router = gin.Default()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	allowOriginFunc := func(origin string) bool { return true }
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "http://localhost:4201"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowOriginFunc:  allowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
	})

	handler := c.Handler(a.Router)

	log.Println("Cardapin Running at http://localhost:" + addr)

	log.Fatal(http.ListenAndServe(addr, handler))
}

func (a *App) initializeRoutes() {
	//public
	a.Router.POST("/login", Login)
	a.Router.GET("/menus/company/:companyCode/enabled", GetMenuEnabledByCompanyCode)

	a.Router.POST("/users", CreateUser)
	a.Router.GET("/users/:id", GetUserById)
	a.Router.POST("/menus", CreateMenu)
	a.Router.PUT("/menus/:id", UpdateMenu)

	a.Router.DELETE("/menus/:id", DeleteMenu)
	a.Router.PUT("/menus/:id/enabled", EnableMenu)
	a.Router.GET("/menus/company", GetMenuByLoggedCompany)

	a.Router.POST("/tables", CreateTable)
	a.Router.PUT("/tables/:id", UpdateTable)
	a.Router.GET("/tables/company/:id", ListTable)
	a.Router.DELETE("/tables/:id", DeleteTable)

	a.Router.POST("/clients", CreateClient)
	a.Router.PUT("/clients/:phone", UpdateClient)
	a.Router.GET("/clients/:phone", GetClientByPhone)

	a.Router.PUT("/companies/:id", UpdateCompany)
	a.Router.GET("/companies/:id", GetCompany)

	a.Router.PATCH("/images", generatePresignedUrlToPut)
	a.Router.GET("/payment-types", GetPaymentTypes)
	a.Router.GET("/sections", GetSections)

	a.Router.POST("/additional-items-groups", CreateAdditionalItemGroup)
	a.Router.PUT("/additional-items-groups/:id", UpdateAdditionalItemGroup)
	a.Router.GET("/additional-items-groups/company/:id", ListAdditionalItemGroup)
	a.Router.DELETE("/additional-items-groups/:id", DeleteAdditionalGroup)

	a.Router.POST("/companies", NewCompanyInterested)

	a.Router.GET("/drop-database", DropDatabase) //JUST FOR LOCAL ENVIRONMENT
}

func Login(c *gin.Context) {
	var userJSON model.UserLoginJSON

	if err := c.ShouldBindJSON(&userJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := &services.UserService{}
	userResponseJSON, err := userService.Login(userJSON.AsModel())
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	token, _ := auth.CreateToken(userResponseJSON)
	c.Header(httphelper.AuthorizationHeader, "Bearer "+token)
	c.Header(httphelper.AccessControlExposeHeaders, httphelper.AuthorizationHeader)
	c.JSON(http.StatusOK, userResponseJSON)
}

func CreateUser(c *gin.Context) {
	var userJSON model.UserRequestJSON

	if err := c.ShouldBindJSON(&userJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := services.UserService{}
	userResponseJSON, err := userService.SaveUser(userJSON.AsModel())
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	httphelper.HandleResponse(userResponseJSON, c.Writer, err)
}

func generatePresignedUrlToPut(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	var fileJson model.FileRequest

	if err := c.ShouldBindJSON(&fileJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageService := services.ImageServices{}
	jsonValue, _ := imageService.GeneratePresignedUrlToPut(fileJson)
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	c.JSON(http.StatusOK, jsonValue)
}

func GetUserById(c *gin.Context) {
	userService := services.UserService{}
	if _, err := auth.TokenValid(c.GetHeader(httphelper.AuthorizationHeader)); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	userResponseJSON, _ := userService.GetUserById(c.Param("id"))
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	if userResponseJSON.ID == 0 {
		c.JSON(http.StatusOK, "{}")
	} else {
		c.JSON(http.StatusOK, userResponseJSON)
	}
}

func CreateMenu(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}

	var menuJson model.MenuJSON

	if err := c.ShouldBindJSON(&menuJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error:" + err.Error())
		return
	}

	menuServices := services.MenuServices{}
	menuObject := *menuJson.AsModel()
	if httphelper.HasConflict(&menuObject, c.Writer) {
		return
	}
	menuJSON, err := menuServices.Save(menuObject)
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	httphelper.HandleResponse(menuJSON, c.Writer, err)
}

func CreateCompany(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	var companyJson model.CompanyJson

	if err := c.ShouldBindJSON(&companyJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyServices := services.CompanyServices{}
	company, err := companyServices.Save(*companyJson.AsModel())
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	companyJSON := company.AsJson()
	httphelper.HandleResponse(companyJSON, c.Writer, err)
}

func UpdateCompany(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	var companyJSONRequest model.CompanyJson

	if err := c.ShouldBindJSON(&companyJSONRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyServices := services.CompanyServices{}
	company, err := companyServices.Update(*companyJSONRequest.AsModel())
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	companyJSON := company.AsJson()
	httphelper.HandleResponse(companyJSON, c.Writer, err)
}

func GetCompany(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	companyServices := services.CompanyServices{}
	company := companyServices.List(c.Param("id"))
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	companyJSON := company.AsJson()
	httphelper.HandleResponse(companyJSON, c.Writer, nil)
}

func GetMenuEnabledByCompanyCode(c *gin.Context) {
	menuServices := services.MenuServices{}
	menuJSON, err := menuServices.GetMenuEnabledByCompanyCode(c.Param("companyCode"))
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	httphelper.HandleResponse(menuJSON, c.Writer, err)
}

func EnableMenu(c *gin.Context) {
	menuServices := services.MenuServices{}
	menuJSON, err := menuServices.EnableMenu(c.Param("id"))
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	httphelper.HandleResponse(menuJSON, c.Writer, err)
}

func GetMenuByLoggedCompany(c *gin.Context) {
	menuServices := services.MenuServices{}

	menuJSON, err := menuServices.GetMenuByLoggedCompany(httphelper.GetToken(c.Request))
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	encodeFunction := func(resp http.ResponseWriter, jsonModel interface{}) {
		json.NewEncoder(resp).Encode(jsonModel)
	}
	httphelper.HandleEmptySliceResp(len(menuJSON), c.Writer, encodeFunction, menuJSON, err)
}

func UpdateMenu(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	var menuJSON model.MenuJSON
	_ = json.NewDecoder(c.Request.Body).Decode(&menuJSON)
	menuServices := services.MenuServices{}
	menuObject := *menuJSON.AsModel()
	if httphelper.HasConflict(&menuObject, c.Writer) {
		return
	}
	menuResponse, err := menuServices.Update(menuObject)
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	httphelper.HandleResponse(menuResponse, c.Writer, err)
}

func DeleteMenu(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	var menuJSON model.MenuJSON
	_ = json.NewDecoder(c.Request.Body).Decode(&menuJSON)
	menuServices := services.MenuServices{}
	if _, err := menuServices.DeleteMenu(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, "{ \"message\": \""+err.Error()+"\"}")
		return
	}
	c.Status(http.StatusNoContent)
}

func CreateClient(c *gin.Context) {
	var client model.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientServices := services.ClientService{}
	savedClient, _ := clientServices.Save(client, false)
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	c.JSON(http.StatusOK, savedClient)
}

func UpdateClient(c *gin.Context) {
	var client model.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientServices := services.ClientService{}
	savedClient, _ := clientServices.Save(client, true)
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	c.JSON(http.StatusOK, savedClient)
}

func GetClientByPhone(c *gin.Context) {

	clientServices := services.ClientService{}
	phone, _ := strconv.ParseUint(c.Param("phone"), 10, 64)
	client, _ := clientServices.GetByPhone(uint64(phone))
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	if client.Phone == 0 {
		c.JSON(http.StatusOK, "{}")
	} else {
		c.JSON(http.StatusOK, client)
	}
}

func isAuthorized(c *gin.Context) bool {
	godotenv.Load()
	disableAuth := os.Getenv("disable_auth")
	if disableAuth == "true" {
		return true
	}
	if _, err := auth.TokenValid(c.Request.Header.Get(httphelper.AuthorizationHeader)); err != nil {
		c.Status(http.StatusUnauthorized)
		return false
	}
	return true
}

func GetPaymentTypes(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	paymentTypeServices := services.PaymentTypeService{}
	paymentTypes, err := paymentTypeServices.ListPaymentTypes()
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	encodeFunction := func(resp http.ResponseWriter, jsonModel interface{}) {
		json.NewEncoder(resp).Encode(jsonModel)
	}
	httphelper.HandleEmptySliceResponse(len(paymentTypes), c.Writer, encodeFunction, paymentTypes, err)
}

func GetSections(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	sectionService := services.SectionService{}
	sections, err := sectionService.ListSection()
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	encodeFunction := func(resp http.ResponseWriter, jsonModel interface{}) {
		json.NewEncoder(resp).Encode(jsonModel)
	}
	httphelper.HandleEmptySliceResponse(len(sections), c.Writer, encodeFunction, sections, err)
}

func CreateTable(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	var tableJSON model.TableJSON

	if err := c.ShouldBindJSON(&tableJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tableServices := services.TableServices{}
	table := tableJSON.AsModel()
	if httphelper.HasConflict(table, c.Writer) {
		return
	}
	savedTable, err := tableServices.Save(*table, false)
	httphelper.HandleResponse(savedTable.AsJSON(), c.Writer, err)
}

func UpdateTable(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	var tableJSON model.TableJSON

	if err := c.ShouldBindJSON(&tableJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tableServices := services.TableServices{}
	table := tableJSON.AsModel()
	savedTable, err := tableServices.Save(*table, true)
	if err != nil {
		httphelper.HandleResponse(nil, c.Writer, err)
		return
	}
	httphelper.HandleResponse(savedTable.AsJSON(), c.Writer, err)
}

func DeleteTable(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	tableServices := services.TableServices{}
	if _, err := tableServices.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, "{ \"message\": \""+err.Error()+"\"}")
		return
	}
	c.Status(http.StatusNoContent)
}

func ListTable(c *gin.Context) {
	tableServices := services.TableServices{}
	tablesJSON, err := tableServices.List(c.Param("id"))
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	encodeFunction := func(resp http.ResponseWriter, jsonModel interface{}) {
		json.NewEncoder(resp).Encode(jsonModel)
	}
	httphelper.HandleEmptySliceResponse(len(tablesJSON), c.Writer, encodeFunction, tablesJSON, err)
}

func CreateAdditionalItemGroup(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	var additionalItemGroupJSON model.AdditionalItemsGroupJSON

	if err := c.ShouldBindJSON(&additionalItemGroupJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	additionalItemGroup := additionalItemGroupJSON.AsModel()
	if httphelper.HasConflict(additionalItemGroup, c.Writer) {
		return
	}
	savedGroup, err := model.SaveAdditionalItemsGroup(*additionalItemGroup)
	httphelper.HandleResponse(savedGroup.AsJSON(), c.Writer, err)
}

func UpdateAdditionalItemGroup(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}
	var additionalItemGroupJSON model.AdditionalItemsGroupJSON

	if err := c.ShouldBindJSON(&additionalItemGroupJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	additionalItemGroup := additionalItemGroupJSON.AsModel()
	if httphelper.HasConflict(additionalItemGroup, c.Writer) {
		return
	}
	savedGroup, err := model.UpdateAdditionalItemsGroup(*additionalItemGroup)
	httphelper.HandleResponse(savedGroup.AsJSON(), c.Writer, err)
}

func ListAdditionalItemGroup(c *gin.Context) {

	groupsJSON := model.ListAdditionalItemsByCompanyId(c.Param("id"))
	c.Header(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	encodeFunction := func(resp http.ResponseWriter, jsonModel interface{}) {
		json.NewEncoder(resp).Encode(jsonModel)
	}
	httphelper.HandleEmptySliceResponse(len(groupsJSON), c.Writer, encodeFunction, groupsJSON, nil)
}

func DeleteAdditionalGroup(c *gin.Context) {
	if ok := isAuthorized(c); !ok {
		return
	}

	err := model.DeleteAdditionalGroup(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "{ \"message\": \""+err.Error()+"\"}")
		return
	}
	c.Status(http.StatusNoContent)
}

func NewCompanyInterested(c *gin.Context) {
	var mailInfo email.Email

	if err := c.ShouldBindJSON(&mailInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyServices := services.CompanyServices{}
	if err := companyServices.CompanyInterested(mailInfo); !err {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func DropDatabase(c *gin.Context) {
	db.DB.Exec("drop schema public cascade")
	db.DB.Exec("create schema public")
	migration.AutoMigrate()
}
