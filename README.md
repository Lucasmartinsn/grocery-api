# grocery-api
In this project we will develop a Rest API in the Golang language.
The API will manage the basic functionalities of a supermarket.


## Requere
### Golang 1.20
### PostgreSQL 15
docker run -d --name banco_grocery_api -p 5432:5432 --network rede-teste -e POSTGRES_PASSWORD=123 -e POSTGRES_DB=grocere_api postgres:15