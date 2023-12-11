curl -X GET localhost:8080/drop-database

curl -X POST localhost:8080/users -H "Content-Type: application/json" -d @request-examples/user-request.json

curl -X PUT localhost:8080/companies/1 -H "Content-Type: application/json" -d @request-examples/company-request.json

curl -X POST localhost:8080/additional-items-groups -H "Content-Type: application/json" -d @request-examples/additional-items-group-request-hamburguers.json

curl -X POST localhost:8080/additional-items-groups -H "Content-Type: application/json" -d @request-examples/additional-items-group-request-bebidas.json

curl -X POST localhost:8080/menus -H "Content-Type: application/json" -d @request-examples/menu-request.json

curl -X POST localhost:8080/clients -H "Content-Type: application/json" -d @request-examples/client-request.json

curl -X POST localhost:8080/users -H "Content-Type: application/json" -d @request-examples/user-request-mala-lanches.json
