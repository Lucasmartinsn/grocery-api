## Information
In this project we will develop a Rest API in the Golang language.
The API will manage the basic functionalities of a supermarket.


## technical requirements
* Golang 1.20
* PostgreSQL 15
* Basic knowledge of Golang and PostgreSQL

## Resquests
* All routes except the login route need an authentication token, send the token in the header
* Must have in the header
    ```
    Authorization: Bearer jhkhkjhkjhkjcxbnc....
    Content-Type: application/json
    ```
### Employee  
``Login employee ``

The ID is the field ID from the t_employee table

 * Url: http://localhost:5000/api/employee/login
 * Method: POST
 * Body
   ```json
        {
            "cpf": 99999999999,
            "password": "root"
        }
    ```
 * Response
    ```json
    {
        "token":"hjksadjahdbmnzbcyudy...."
    } 
    ```  

``Search employee `` 

To search for employees, we have several routes. 
* Get all employee
    *  Url: http://localhost:5000/api/employee/
* Get one employee
    *  Url: http://localhost:5000/api/employee/?id=2342
* Get all employees with active or false status
    *  Url: http://localhost:5000/api/employee/?active=true
* Get one employee with active and id
    *  Url: http://localhost:5000/api/employee/?active=true&id=12313

 * Method: GET
 * Response
    ```json
    {
        "employee": [
            {
                "id": "bafb12d5-062b-4fc2-b4d0-456c9eaffee9",
                "name": "admin",
                "cpf": 99999999999,
                "password": "123",
                "office": "ceo",
                "active": true,
                "admin": true,
                "createon_date": "2024-05-31T11:50:00.038517Z"
            }
        ]
    } 
    ```  
``Create employee ``
 * Url: http://localhost:5000/api/employee/
 * Method: POST
 * Body
   ```json
    {
        "name":"admin",
        "cpf":88888888888,
        "password":"123",
        "office":"ceo",
        "active": true,
        "admin":true
    }
    ```
 ``Delete employee registration``
 * Url: http://localhost:5000/api/employee/ID
 * Method: DELETE
 * Body
   ```json
    {
        "name":"admin",
        "cpf":88888888888,
        "password":"123",
        "office":"ceo",
        "active": true,
        "admin":true
    }
    ```
``Update employee registration`` 

To update for employees, we have several routes. 
* Update all employee fields
    *  Url: http://localhost:5000/api/employee/ID?all=true
        * Method: PUT
        * Body
            ```json
            {
                "name": "admin",
                "password": "123",
                "office": "ceo",
                "active": true,
                "admin": true,
            } 
            ```
* Update one employee field
    *  Url: http://localhost:5000/api/employee/ID?pass=true
        * Method: PUT
        * Body
            ```json
            {
                "password": "123"
            } 
            ```
* Update one employee field
    *  Url: http://localhost:5000/api/employee/ID?name=true
        * Method: PUT
        * Body
            ```json
            {
                "name": "admin"
            } 
            ```
* Update one employee field
    *  Url: http://localhost:5000/api/employee/ID?office=true
        * Method: PUT
        * Body
            ```json
            {
                "office": "ceo"
            } 
            ```
* Update one employee field
    *  Url: http://localhost:5000/api/employee/ID?active=true
        * Method: PUT
        * Body
            ```json
            {
                "active": true
            } 
            ```
* Update one employee field
    *  Url: http://localhost:5000/api/employee/ID?admin=true
        * Method: PUT
        * Body
            ```json
            {
                "admin": true
            } 
            ```