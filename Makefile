.PHONY: clean build deploy

build:
	mkdir ./bin
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/client/createNewClient handlers/client/createNewClient/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/client/updateclient handlers/client/updateClient/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/client/getClientByPhone handlers/client/getClientByPhone/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/company/updateCompanyHandler handlers/company/updateCompanyHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/company/companyInterested handlers/company/companyInterested/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/image/imagePatchHandler handlers/image/imagePatchHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/menu/createMenuHandler handlers/menu/createMenuHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/menu/listAllMenuHandler handlers/menu/listAllMenuHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/menu/listEnabledMenuHandler handlers/menu/listEnabledMenuHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/menu/enableMenuHandler handlers/menu/enableMenuHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/menu/deleteMenuHandler handlers/menu/deleteMenuHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/menu/updateMenuHandler handlers/menu/updateMenuHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/user/createUserHandler handlers/user/createUserHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/user/getUserHandler handlers/user/getUserHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/user/loginHandler handlers/user/loginHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/user/authorizer handlers/user/authorizer/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/configuration/listPaymentTypesHandler handlers/configuration/listPaymentTypesHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/configuration/listSectionsHandler handlers/configuration/listSectionsHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/table/createTableHandler handlers/table/createTableHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/table/deleteTableHandler handlers/table/deleteTableHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/table/listTableHandler handlers/table/listTableHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/table/updateTableHandler handlers/table/updateTableHandler/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/group/createGroup handlers/group/createGroup/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/group/updateGroup handlers/group/updateGroup/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/group/listGroup handlers/group/listGroup/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/group/deleteGroup handlers/group/deleteGroup/main.go

	zip -vr bin/client/createNewClient.zip bin/client/createNewClient
	zip -vr bin/client/updateClient.zip bin/client/updateClient
	zip -vr bin/client/getClientByPhone.zip bin/client/getClientByPhone
	zip -vr bin/company/updateCompanyHandler.zip bin/company/updateCompanyHandler
	zip -vr bin/company/companyInterested.zip bin/company/companyInterested
	zip -vr bin/image/imagePatchHandler.zip bin/image/imagePatchHandler
	zip -vr bin/menu/createMenuHandler.zip bin/menu/createMenuHandler
	zip -vr bin/menu/listAllMenuHandler.zip bin/menu/listAllMenuHandler
	zip -vr bin/menu/listEnabledMenuHandler.zip bin/menu/listEnabledMenuHandler
	zip -vr bin/menu/enableMenuHandler.zip bin/menu/enableMenuHandler
	zip -vr bin/menu/deleteMenuHandler.zip bin/menu/deleteMenuHandler
	zip -vr bin/menu/updateMenuHandler.zip bin/menu/updateMenuHandler
	zip -vr bin/user/createUserHandler.zip bin/user/createUserHandler
	zip -vr bin/user/getUserHandler.zip bin/user/getUserHandler
	zip -vr bin/user/loginHandler.zip bin/user/loginHandler
	zip -vr bin/user/authorizer.zip bin/user/authorizer
	zip -vr bin/configuration/listPaymentTypesHandler.zip bin/configuration/listPaymentTypesHandler
	zip -vr bin/configuration/listSectionsHandler.zip bin/configuration/listSectionsHandler
	zip -vr bin/table/createTableHandler.zip bin/table/createTableHandler
	zip -vr bin/table/deleteTableHandler.zip bin/table/deleteTableHandler
	zip -vr bin/table/listTableHandler.zip bin/table/listTableHandler
	zip -vr bin/table/updateTableHandler.zip bin/table/updateTableHandler
	zip -vr bin/group/createGroup.zip bin/group/createGroup
	zip -vr bin/group/updateGroup.zip bin/group/updateGroup
	zip -vr bin/group/listGroup.zip bin/group/listGroup
	zip -vr bin/group/deleteGroup.zip bin/group/deleteGroup
	
clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
