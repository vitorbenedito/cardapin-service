package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jinzhu/gorm"

	"cardap.in/lambda/db"
	"cardap.in/lambda/migration"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nsf/jsondiff"
)

const (
	POST   = "POST"
	PUT    = "PUT"
	GET    = "GET"
	DELETE = "DELETE"
)

var a App
var cleaner func()

func TestMain(m *testing.M) {
	createDatabase()
	migration.AutoMigrate()
	a.Initialize()
	cleaner = dropDatabaseTest(db.DB)
	code := m.Run()
	os.Exit(code)
}

func TestGetPaymentTypes(t *testing.T) {
	defer cleaner()
	response := executeRequest(request(GET, "/payment-types", ""))
	checkResponseCode(t, http.StatusOK, response.Code)
	validate(response, t, "expected-responses/get-payment-types.json")
}

func TestCreateUser(t *testing.T) {
	defer cleaner()
	response := executeRequest(request(POST, "/users", "requests/user-request.json"))
	checkResponseCode(t, http.StatusOK, response.Code)
	validate(response, t, "expected-responses/post-user.json")
}

func TestCompanyInterested(t *testing.T) {
	defer cleaner()
	response := executeRequest(request(POST, "/companies", "requests/company-interested.json"))
	checkResponseCode(t, http.StatusOK, response.Code)
	validate(response, t, "expected-responses/post-company.json")
}
func TestCreateCompanies(t *testing.T) {
	defer cleaner()
	executeRequest(request(POST, "/users", "requests/user-request.json"))
	response := executeRequest(request(PUT, "/companies/1", "requests/company-request.json"))
	checkResponseCode(t, http.StatusOK, response.Code)
	validate(response, t, "expected-responses/put-company.json")
}

func TestCreateCompaniesAndRemoveAddress(t *testing.T) {
	defer cleaner()
	executeRequest(request(POST, "/users", "requests/user-request.json"))

	response := executeRequest(request(PUT, "/companies/1", "requests/company-request.json"))
	checkResponseCode(t, http.StatusOK, response.Code)
	validate(response, t, "expected-responses/put-company.json")

	responseRemoveAddress := executeRequest(request(PUT, "/companies/1", "requests/company-request-delete-address.json"))
	checkResponseCode(t, http.StatusOK, responseRemoveAddress.Code)
	validate(responseRemoveAddress, t, "expected-responses/put-company-remove-address.json")

	getCompanyResp := executeRequest(request(GET, "/companies/1", ""))

	validate(getCompanyResp, t, "expected-responses/put-company-remove-address.json")
}

func TestCreateAdditionalItemsGroups(t *testing.T) {
	defer cleaner()
	executeRequest(request(POST, "/users", "requests/user-request.json"))
	response := executeRequest(request(POST, "/additional-items-groups", "requests/additional-items-group-request-hamburguers.json"))
	checkResponseCode(t, http.StatusOK, response.Code)
	validate(response, t, "expected-responses/post-additional-items-group-hamburguers.json")

	response = executeRequest(request(POST, "/additional-items-groups", "requests/additional-items-group-request-bebidas.json"))
	checkResponseCode(t, http.StatusOK, response.Code)
	validate(response, t, "expected-responses/post-additional-items-group-bebidas.json")
}

func TestCreateMenus(t *testing.T) {
	defer cleaner()
	executeRequest(request(POST, "/users", "requests/user-request.json"))
	response := executeRequest(request(POST, "/menus", "requests/menu-request.json"))
	checkResponseCode(t, http.StatusOK, response.Code)
	validate(response, t, "expected-responses/post-menu.json")
}

func TestCreateMenusWithDuplicatedCateogry(t *testing.T) {
	defer cleaner()
	executeRequest(request(POST, "/users", "requests/user-request.json"))
	response := executeRequest(request(POST, "/menus", "requests/menu-request-same-category.json"))
	checkResponseCode(t, http.StatusConflict, response.Code)
}

func TestCreateMenusWithDuplicatedProduct(t *testing.T) {
	defer cleaner()
	executeRequest(request(POST, "/users", "requests/user-request.json"))
	response := executeRequest(request(POST, "/menus", "requests/menu-request-same-product.json"))
	checkResponseCode(t, http.StatusConflict, response.Code)
}

func TestCreateClients(t *testing.T) {
	defer cleaner()
	response := executeRequest(request(POST, "/clients", "requests/client-request.json"))
	checkResponseCode(t, http.StatusOK, response.Code)
	validate(response, t, "expected-responses/post-client.json")
}

func request(method string, endpoint string, file string) *http.Request {
	var buff *bytes.Buffer
	if file != "" {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		if content != nil {
			buff = bytes.NewBuffer(content)
		}
	}
	if buff != (*bytes.Buffer)(nil) {
		req, err := http.NewRequest(method, endpoint, buff)
		if err != nil {
			log.Fatal(err)
		}
		return req
	}
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	return req
}

func validate(resp *httptest.ResponseRecorder, t *testing.T, file string) {
	opts := jsondiff.DefaultConsoleOptions()
	body := resp.Body.String()
	expected, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	r, diff := jsondiff.Compare([]byte(body), expected, &opts)
	if r == jsondiff.NoMatch {
		t.Errorf("Expected an json with different values. \r\n\r\nGot:\r\n\r\n%s\r\n\r\nExpected:\r\n\r\n%s\r\n\r\nDiff:\r\n\r\n%s", body, string(expected), diff)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func createDatabase() {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)

	database := "cardappin"
	// pulls an image, creates a container based on it and runs it
	var err error
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "localhost", "5433", "test", database, "test")
	db2, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Error to open connection: %s", err)
	}
	c := db.Connector{}
	// db2.LogMode(true)
	c.InitializeDatabaseParam(db2)
}

func dropDatabaseTest(db *gorm.DB) func() {
	return func() {
		db.Exec("drop schema public cascade")
		db.Exec("create schema public")
		migration.AutoMigrate()
	}
}
