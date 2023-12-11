package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"cardap.in/lambda/auth"
	"cardap.in/lambda/db"
	"cardap.in/lambda/email"
	"cardap.in/lambda/httphelper"
	"cardap.in/lambda/migration"
	"cardap.in/lambda/model"
	"cardap.in/lambda/services"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()

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
	a.Router.HandleFunc("/login", Login).Methods("POST")
	a.Router.HandleFunc("/menus/company/{companyCode}/enabled", GetMenuEnabledByCompanyCode).Methods("GET")

	a.Router.HandleFunc("/users", CreateUser).Methods("POST")
	a.Router.HandleFunc("/users/{id}", GetUserById).Methods("GET")
	a.Router.HandleFunc("/menus", CreateMenu).Methods("POST")
	a.Router.HandleFunc("/menus/{id}", UpdateMenu).Methods("PUT")

	a.Router.HandleFunc("/menus/{id}", DeleteMenu).Methods("DELETE")
	a.Router.HandleFunc("/menus/{id}/enabled", EnableMenu).Methods("PUT")
	a.Router.HandleFunc("/menus/company", GetMenuByLoggedCompany).Methods("GET")

	a.Router.HandleFunc("/tables", CreateTable).Methods("POST")
	a.Router.HandleFunc("/tables/{id}", UpdateTable).Methods("PUT")
	a.Router.HandleFunc("/tables/company/{id}", ListTable).Methods("GET")
	a.Router.HandleFunc("/tables/{id}", DeleteTable).Methods("DELETE")

	a.Router.HandleFunc("/clients", CreateClient).Methods("POST")
	a.Router.HandleFunc("/clients/{phone}", UpdateClient).Methods("PUT")
	a.Router.HandleFunc("/clients/{phone}", GetClientByPhone).Methods("GET")

	a.Router.HandleFunc("/companies/{id}", UpdateCompany).Methods("PUT")

	a.Router.HandleFunc("/images", generatePresignedUrlToPut).Methods("PATCH")
	a.Router.HandleFunc("/payment-types", GetPaymentTypes).Methods("GET")
	a.Router.HandleFunc("/sections", GetSections).Methods("GET")

	a.Router.HandleFunc("/additional-items-groups", CreateAdditionalItemGroup).Methods("POST")
	a.Router.HandleFunc("/additional-items-groups/{id}", UpdateAdditionalItemGroup).Methods("PUT")
	a.Router.HandleFunc("/additional-items-groups/company/{id}", ListAdditionalItemGroup).Methods("GET")
	a.Router.HandleFunc("/additional-items-groups/{id}", DeleteAdditionalGroup).Methods("DELETE")

	a.Router.HandleFunc("/companies/{id}", UpdateCompany).Methods("PUT")

	a.Router.HandleFunc("/companies", NewCompanyInterested).Methods("POST")

	a.Router.HandleFunc("/drop-database", DropDatabase).Methods("GET") //JUST FOR LOCAL ENVIRONMENT
}

func Login(resp http.ResponseWriter, req *http.Request) {
	var userJSON model.UserLoginJSON
	_ = json.NewDecoder(req.Body).Decode(&userJSON)
	userService := &services.UserService{}
	userResponseJSON, err := userService.Login(userJSON.AsModel())
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	token, _ := auth.CreateToken(userResponseJSON)
	resp.Header().Add(httphelper.AuthorizationHeader, "Bearer "+token)
	resp.Header().Add(httphelper.AccessControlExposeHeaders, httphelper.AuthorizationHeader)
	json.NewEncoder(resp).Encode(userResponseJSON)
}

func CreateUser(resp http.ResponseWriter, req *http.Request) {
	var userJSON model.UserRequestJSON
	_ = json.NewDecoder(req.Body).Decode(&userJSON)
	userService := services.UserService{}
	userResponseJSON, err := userService.SaveUser(userJSON.AsModel())
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	httphelper.HandleResponse(userResponseJSON, resp, err)
}

func generatePresignedUrlToPut(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	var fileJson model.FileRequest
	_ = json.NewDecoder(req.Body).Decode(&fileJson)
	imageService := services.ImageServices{}
	jsonValue, _ := imageService.GeneratePresignedUrlToPut(fileJson)
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	resp.Write([]byte(jsonValue))
}

func GetUserById(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userService := services.UserService{}
	if _, err := auth.TokenValid(req.Header.Get(httphelper.AuthorizationHeader)); err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}
	userResponseJSON, _ := userService.GetUserById(params["id"])
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	if userResponseJSON.ID == 0 {
		resp.Write([]byte("{}"))
	} else {
		json.NewEncoder(resp).Encode(userResponseJSON)
	}
}

func CreateMenu(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	var menuJson model.MenuJSON
	json.NewDecoder(req.Body).Decode(&menuJson)
	menuServices := services.MenuServices{}
	menuObject := *menuJson.AsModel()
	if httphelper.HasConflict(&menuObject, resp) {
		return
	}
	menuJSON, err := menuServices.Save(menuObject)
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	httphelper.HandleResponse(menuJSON, resp, err)
}

func CreateCompany(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	var companyJson model.CompanyJson
	_ = json.NewDecoder(req.Body).Decode(&companyJson)
	companyServices := services.CompanyServices{}
	company, err := companyServices.Save(*companyJson.AsModel())
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	companyJSON := company.AsJson()
	httphelper.HandleResponse(companyJSON, resp, err)
}

func UpdateCompany(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	var companyJSONRequest model.CompanyJson
	_ = json.NewDecoder(req.Body).Decode(&companyJSONRequest)
	companyServices := services.CompanyServices{}
	company, err := companyServices.Update(*companyJSONRequest.AsModel())
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	companyJSON := company.AsJson()
	httphelper.HandleResponse(companyJSON, resp, err)
}

func GetMenuEnabledByCompanyCode(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	menuServices := services.MenuServices{}
	menuJSON, err := menuServices.GetMenuEnabledByCompanyCode(params["companyCode"])
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	httphelper.HandleResponse(menuJSON, resp, err)
}

func EnableMenu(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	menuServices := services.MenuServices{}
	menuJSON, err := menuServices.EnableMenu(params["id"])
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	httphelper.HandleResponse(menuJSON, resp, err)
}

func GetMenuByLoggedCompany(resp http.ResponseWriter, req *http.Request) {
	menuServices := services.MenuServices{}

	menuJSON, err := menuServices.GetMenuByLoggedCompany(httphelper.GetToken(req))
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	encodeFunction := func(resp http.ResponseWriter, jsonModel interface{}) {
		json.NewEncoder(resp).Encode(jsonModel)
	}
	httphelper.HandleEmptySliceResp(len(menuJSON), resp, encodeFunction, menuJSON, err)
}

func UpdateMenu(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	var menuJSON model.MenuJSON
	_ = json.NewDecoder(req.Body).Decode(&menuJSON)
	menuServices := services.MenuServices{}
	menuObject := *menuJSON.AsModel()
	if httphelper.HasConflict(&menuObject, resp) {
		return
	}
	menuResponse, err := menuServices.Update(menuObject)
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	httphelper.HandleResponse(menuResponse, resp, err)
}

func DeleteMenu(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	var menuJSON model.MenuJSON
	_ = json.NewDecoder(req.Body).Decode(&menuJSON)
	menuServices := services.MenuServices{}
	if _, err := menuServices.DeleteMenu(mux.Vars(req)["id"]); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("{ \"message\": \"" + err.Error() + "\"}"))
		return
	}
	resp.WriteHeader(http.StatusNoContent)
}

func CreateClient(resp http.ResponseWriter, req *http.Request) {
	var client model.Client
	_ = json.NewDecoder(req.Body).Decode(&client)
	clientServices := services.ClientService{}
	savedClient, _ := clientServices.Save(client, false)
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	json.NewEncoder(resp).Encode(savedClient)
}

func UpdateClient(resp http.ResponseWriter, req *http.Request) {
	var client model.Client
	_ = json.NewDecoder(req.Body).Decode(&client)
	clientServices := services.ClientService{}
	savedClient, _ := clientServices.Save(client, true)
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	json.NewEncoder(resp).Encode(savedClient)
}

func GetClientByPhone(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	clientServices := services.ClientService{}
	phone, _ := strconv.ParseUint(params["phone"], 10, 64)
	client, _ := clientServices.GetByPhone(uint64(phone))
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	if client.Phone == 0 {
		resp.Write([]byte("{}"))
	} else {
		json.NewEncoder(resp).Encode(client)
	}
}

func isAuthorized(resp http.ResponseWriter, req *http.Request) bool {
	godotenv.Load()
	disableAuth := os.Getenv("disable_auth")
	if disableAuth == "true" {
		return true
	}
	if _, err := auth.TokenValid(req.Header.Get(httphelper.AuthorizationHeader)); err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}

func GetPaymentTypes(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	paymentTypeServices := services.PaymentTypeService{}
	paymentTypes, err := paymentTypeServices.ListPaymentTypes()
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	encodeFunction := func(resp http.ResponseWriter, jsonModel interface{}) {
		json.NewEncoder(resp).Encode(jsonModel)
	}
	httphelper.HandleEmptySliceResponse(len(paymentTypes), resp, encodeFunction, paymentTypes, err)
}

func GetSections(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	sectionService := services.SectionService{}
	sections, err := sectionService.ListSection()
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	encodeFunction := func(resp http.ResponseWriter, jsonModel interface{}) {
		json.NewEncoder(resp).Encode(jsonModel)
	}
	httphelper.HandleEmptySliceResponse(len(sections), resp, encodeFunction, sections, err)
}

func CreateTable(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	var tableJSON model.TableJSON
	_ = json.NewDecoder(req.Body).Decode(&tableJSON)
	tableServices := services.TableServices{}
	table := tableJSON.AsModel()
	if httphelper.HasConflict(table, resp) {
		return
	}
	savedTable, err := tableServices.Save(*table, false)
	httphelper.HandleResponse(savedTable.AsJSON(), resp, err)
}

func UpdateTable(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	var tableJSON model.TableJSON
	_ = json.NewDecoder(req.Body).Decode(&tableJSON)
	tableServices := services.TableServices{}
	table := tableJSON.AsModel()
	savedTable, err := tableServices.Save(*table, true)
	if err != nil {
		httphelper.HandleResponse(nil, resp, err)
		return
	}
	httphelper.HandleResponse(savedTable.AsJSON(), resp, err)
}

func DeleteTable(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	tableServices := services.TableServices{}
	if _, err := tableServices.Delete(mux.Vars(req)["id"]); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("{ \"message\": \"" + err.Error() + "\"}"))
		return
	}
	resp.WriteHeader(http.StatusNoContent)
}

func ListTable(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	tableServices := services.TableServices{}
	tablesJSON, err := tableServices.List(params["id"])
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	encodeFunction := func(resp http.ResponseWriter, jsonModel interface{}) {
		json.NewEncoder(resp).Encode(jsonModel)
	}
	httphelper.HandleEmptySliceResponse(len(tablesJSON), resp, encodeFunction, tablesJSON, err)
}

func CreateAdditionalItemGroup(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	var additionalItemGroupJSON model.AdditionalItemsGroupJSON
	_ = json.NewDecoder(req.Body).Decode(&additionalItemGroupJSON)
	additionalItemGroup := additionalItemGroupJSON.AsModel()
	if httphelper.HasConflict(additionalItemGroup, resp) {
		return
	}
	savedGroup, err := model.SaveAdditionalItemsGroup(*additionalItemGroup)
	httphelper.HandleResponse(savedGroup.AsJSON(), resp, err)
}

func UpdateAdditionalItemGroup(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	var additionalItemGroupJSON model.AdditionalItemsGroupJSON
	_ = json.NewDecoder(req.Body).Decode(&additionalItemGroupJSON)
	additionalItemGroup := additionalItemGroupJSON.AsModel()
	if httphelper.HasConflict(additionalItemGroup, resp) {
		return
	}
	savedGroup, err := model.UpdateAdditionalItemsGroup(*additionalItemGroup)
	httphelper.HandleResponse(savedGroup.AsJSON(), resp, err)
}

func ListAdditionalItemGroup(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	groupsJSON := model.ListAdditionalItemsByCompanyId(params["id"])
	resp.Header().Add(httphelper.ContentTypeHeader, httphelper.ApplicationJSONValue)
	encodeFunction := func(resp http.ResponseWriter, jsonModel interface{}) {
		json.NewEncoder(resp).Encode(jsonModel)
	}
	httphelper.HandleEmptySliceResponse(len(groupsJSON), resp, encodeFunction, groupsJSON, nil)
}

func DeleteAdditionalGroup(resp http.ResponseWriter, req *http.Request) {
	if ok := isAuthorized(resp, req); !ok {
		return
	}
	params := mux.Vars(req)
	err := model.DeleteAdditionalGroup(params["id"])
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("{ \"message\": \"" + err.Error() + "\"}"))
		return
	}
	resp.WriteHeader(http.StatusNoContent)
}

func NewCompanyInterested(resp http.ResponseWriter, req *http.Request) {
	var mailInfo email.Email
	_ = json.NewDecoder(req.Body).Decode(&mailInfo)
	companyServices := services.CompanyServices{}
	if err := companyServices.CompanyInterested(mailInfo); !err {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.WriteHeader(http.StatusOK)
}

func DropDatabase(resp http.ResponseWriter, req *http.Request) {
	db.DB.Exec("drop schema public cascade")
	db.DB.Exec("create schema public")
	migration.AutoMigrate()
}
