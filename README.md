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
* Data exchange between server and client must be encrypted. Where AES encryption is used
    * Example of sending and receiving
        * Sending
            ```json
            {
                "data": "sjdgdfjgkldvbmxbmnvbuioeg...",
            }
            ```
        * Receiving
        * Upon receiving data, the name of the ‘data’ key will change according to the route being called
            ```json
            {
                "data": "sjdgdfjgkldvbmxbmnvbuioeg...",
            }
            ```
## Client-side usage
example of using the client-side encryption key in a React application
### Requirements
* Framework React 
    * required packages
        *  crypto-js
        *  dotenv

* Example file .js
    ```javascript
        // ./Encrypt.js
        import CryptoJS from 'crypto-js';
        // returns the value of the .env environment variable
        const key = CryptoJS.enc.Utf8.parse(process.env.REACT_APP_API_KEY);

        export function DecryptData(encryptedData) {
            // Decode Base64 string to bytes
            const encryptedBytes = CryptoJS.enc.Base64.parse(encryptedData);

            // Separate the IV from the encrypted bytes
            const iv = CryptoJS.lib.WordArray.create(encryptedBytes.words.  slice(0, 4), 16);
            const ciphertext = CryptoJS.lib.WordArray.create(encryptedBytes.words.slice(4), encryptedBytes.sigBytes - 16);

        // Decrypt using key and IV
        const decrypted = CryptoJS.AES.decrypt(
                { ciphertext: ciphertext },
                key,
                {
                    iv: iv,
                    mode: CryptoJS.mode.CBC,
                    padding: CryptoJS.pad.Pkcs7
                }
            );

            // Return decrypted data as UTF-8 text
            return JSON.parse(CryptoJS.enc.Utf8.stringify(decrypted));
        };

        export function EncryptData(data) {
            const iv = CryptoJS.lib.WordArray.random(16);  // Gerando IV (Initialization Vector) aleatório
            const encrypted = CryptoJS.AES.encrypt(JSON.stringify(data), key, {
                iv: iv,
                mode: CryptoJS.mode.CBC,
                padding: CryptoJS.pad.Pkcs7
            });

            // Returning IV and concatenated encrypted data
        return iv.concat(encrypted.ciphertext).toString(CryptoJS.enc.Base64);
        };
    ```


## Routes
Here you will see all the routes available for use in the app.
### Employee  
``Login employee ``

The ID is the field ID from the t_employee table

 * Url: http://localhost:5000/api/employee/login
 * Method: POST
 * Body:
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
 * Body:
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
 * Body:
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
        * Body:
            ```json
            {
                "name": "admin",
                "password": "123",
                "office": "ceo",
                "active": true,
                "admin": true,
            } 
            ```
* Update one employee passaword field
    *  Url: http://localhost:5000/api/employee/ID?pass=true
        * Method: PUT
        * Body:
            ```json
            {
                "password": "123"
            } 
            ```
* Update one employee name field
    *  Url: http://localhost:5000/api/employee/ID?name=true
        * Method: PUT
        * Body:
            ```json
            {
                "name": "admin"
            } 
            ```
* Update one employee office field
    *  Url: http://localhost:5000/api/employee/ID?office=true
        * Method: PUT
        * Body:
            ```json
            {
                "office": "ceo"
            } 
            ```
* Update one employee active field
    *  Url: http://localhost:5000/api/employee/ID?active=true
        * Method: PUT
        * Body:
            ```json
            {
                "active": true
            } 
            ```
* Update the employee admin field
    *  Url: http://localhost:5000/api/employee/ID?admin=true
        * Method: PUT
        * Body:
            ```json
            {
                "admin": true
            } 
            ```
### Supplier  
``Search supplier ``
* This route will return all Supplier records
    * Url: http://localhost:5000/api/supplier/
    * Method: GET
    * Response
        ```json
        {
            "supplier": [
                {
                    "id": "e5c4a8f4-c55a-4dbd-b706-ef1759a38f50",
                    "name": "supplier name",
                    "cnpj": 2334455677,
                    "contract_number": 2233445566,
                    "company_name": "supplear company name",
                    "status": true
                }
            ]
        } 
        ```  
* This route will return all products linked to a Supplier,  in the url insert the supplier's ID
    * Url: http://localhost:5000/api/supplier/product/ID
    * Method: GET
    * Response
        ```json
            {
                "Supplier": {
                    "id": "e5c4a8f4-c55a-4dbd-b706-ef1759a38f50",
                    "name": "supplier name",
                    "cnpj": 2334455677,
                    "contract_number": 2233445566,
                    "company_name": "supplear company name",
                    "status": true
                },
                "Products": [
                    {
                        "id": "bee355c9-d441-4587-afea-65b7b28703a7",
                        "batch_id": "a7ce9b17-f30e-4231-bff1-bcbfd5430d69",
                        "supplier_id": "e5c4a8f4-c55a-4dbd-b706-ef1759a38f50",
                        "name": "chease",
                        "volume": 480,
                        "unit_price": 20,
                        "validity": "0001-01-01T00:00:00Z"
                    }
                ]
            }
        ```  
* This route will return all batchs linked to a Supplier,  in the url insert the supplier's ID
    * Url: http://localhost:5000/api/supplier/batch/ID
    * Method: GET
    * Response
        ```json
            {
                "Supplier": {
                    "id": "e5c4a8f4-c55a-4dbd-b706-ef1759a38f50",
                    "name": "supplier name",
                    "cnpj": 2334455677,
                    "contract_number": 2233445566,
                    "company_name": "supplear company name",
                    "status": true
                },
                "Batchs": [
                    {
                        "id": "9ee1b7d5-c74f-4d0d-93e2-3dd4150d9ac2",
                        "supplier_id": "e5c4a8f4-c55a-4dbd-b706-ef1759a38f50",
                        "volume": 500,
                        "price": 5000,
                        "purchase_date": "2024-06-03T17:55:57.353636Z",
                        "delivery_date": "2024-06-30T09:30:30Z"
                    }
                ]
            }
        ```  
``Create supplier ``
* Url: http://localhost:5000/api/supplier/?supplier=true
* Method: POST
* Body:
   ```json
    {
        "name": "supplier name",
        "cnpj":2334455677,
        "contract_number":2233445566,
        "company_name": "supplear company name",
        "status": true
    }
    ```
``Create batch from a supplier ``
* Url: http://localhost:5000/api/supplier/?batch=true
* Method: POST
* Body:
   ```json
    {
        "supplier_id": "e5c4a8f4-c55a-4dbd-b706-ef1759a38f50",
        "volume": 500,
        "price":5000,
        "delivery_date":"2024-06-30T09:30:30Z"
    }
    ```
``Create product from a supplier ``
* Url: http://localhost:5000/api/supplier/?product=true
* Method: POST
* Body:
   ```json
    {
        "batch_id":"4d5e4d6f-2721-42b1-9dba-966c2413e10b",
        "supplier_id": "692f43b6-be73-4c2f-a5e2-cff87302a9ae",
        "name":"ham",
        "volume":500,
        "unit_price":15.50,
        "validity":"2024-07-30T09:30:30Z"
    }
    ```
``Update a vendor's product``
* Url: http://localhost:5000/api/supplier/ID?product=true
* Method: PUT
* Body:
   ```json
    {
        "name":"chease",
        "volume":500,
        "unit_price":20.00
    }
    ```
``Update product volume from a supplier ``
* Url: http://localhost:5000/api/supplier/ID?product=true&volume=true
* Method: PUT
* Body:
   ```json
    {
        "volume":500
    }
    ```
``update supplier data``
* Url: http://localhost:5000/api/supplier/ID?supplier=true
* Method: PUT
* Body:
   ```json
    {
        "name": "supplier name",
        "cnpj":2334455677,
        "contract_number":2233445566,
        "company_name": "supplear company name",
        "status": true
    }
    ```
``update batch data from a supplier``
* Url: http://localhost:5000/api/supplier/ID?batch=true
* Method: PUT
* Body:
   ```json
    {
        "volume": 520,
        "price":5500,
        "delivery_date":"2024-06-30T09:30:30Z"
    }
    ```
### Customer
In customer routes, we use a new token and new encryption to save passwords.
Token submission remains the same as for other routes, but the client token is only accepted on client routes, and so on.

``Login Customer``
* Url: http://localhost:5000/api/customer/login/
* Method: POST
* Body:
   ```json
    {
        "cpf":1234,
        "password":"123"
    }
    ```
* Resoonse
    ```json
    {
        "token":"hjksadjahdbmnzbcyudy...."
    } 
    ```
``Login Employee``
* Url: http://localhost:5000/api/employee/login/customer
* Method: POST
* Body:
   ```json
    {
        "cpf":1234,
        "password":"123"
    }
    ```
* Resoonse
    ```json
    {
        "token":"hjksadjahdbmnzbcyudy...."
    } 
    ```

``Search Customer``
* This route will return all Customer records
    * Url: http://localhost:5000/api/customer/
    * Method: GET
    * Response
        ```json
        {
            "customer": [
                {
                    "id": "d80373f6-8f4b-40bf-9e04-eb140163fc53",
                    "name": "teste",
                    "email": "teste@gmail",
                    "contact": 1234,
                    "cpf": 1234,
                    "creation_date": "2024-06-04T17:59:17.573315Z"
                }
            ]
        } 
        ``` 
* This route will return one Customer record
    * Url: http://localhost:5000/api/customer/?id=ID
    * Method: GET
    * Response
        ```json
        {
            "customer": [
                {
                    "id": "d80373f6-8f4b-40bf-9e04-eb140163fc53",
                    "name": "teste",
                    "email": "teste@gmail",
                    "contact": 1234,
                    "cpf": 1234,
                    "creation_date": "2024-06-04T17:59:17.573315Z"
                }
            ]
        } 
        ``` 
* This route will return all address linked to a Customer, in the url insert the customer's ID
    * Url: http://localhost:5000/api/customer/address/ID
    * Method: GET
    * Response
        ```json
        {
            "Customer": {
                "id": "d80373f6-8f4b-40bf-9e04-eb140163fc53",
                "name": "",
                "email": "lucas@gmail",
                "contact": 1234,
                "cpf": 1234,
                "password": "",
                "creation_date": "2024-06-04T17:59:17.573315Z"
            },
            "Address": [
                {
                    "id": "6a684036-ae25-4341-9d0a-235ff4247a3c",
                    "customer_id": "d80373f6-8f4b-40bf-9e04-eb140163fc53",
                    "street": "",
                    "block": "angel",
                    "number": 234,
                    "state": "EUA"
                }
            ]
        }
        ``` 
* This route will return all Credit cards linked to a Customer, in the url insert the customer's ID
    * Url: http://localhost:5000/api/customer/card/ID
    * Method: GET
    * Response
        ```json
        {
            "Customer": {
                "id": "d80373f6-8f4b-40bf-9e04-eb140163fc53",
                "name": "",
                "email": "lucas@gmail",
                "contact": 1234,
                "cpf": 1234,
                "password": "",
                "creation_date": "2024-06-04T17:59:17.573315Z"
            },
            "Cards": [
                {
                    "id": "7059579e-b022-4e9d-95af-d1b454fce9ef",
                    "customer_id": "d80373f6-8f4b-40bf-9e04-eb140163fc53",
                    "number": 12343,
                    "csv": 133,
                    "name_card": "teste card",
                    "validity": "08/34"
                }
            ]
        }
        ```  
``Create Customer ``
* Url: http://localhost:5000/api/customer/?customer=true
* Method: POST
* Body:
   ```json
    {
        "name":"teste",
        "email":"teste@gmail",
        "contact":1234,
        "cpf":123,
        "password":"123"
    }
    ```
``Create batch from a supplier ``
* Url: http://localhost:5000/api/customer/?address=true
* Method: POST
* Body:
   ```json
    {
        "customer_id":"3f122513-757d-4ba0-b4dd-4f375b1344d9",
        "street":"street nine",
        "block":"angel",
        "number":234,
        "state":"EUA"
    }
    ```
``Create batch from a supplier ``
* Url: http://localhost:5000/api/customer/?card=true
* Method: POST
* Body:
   ```json
    {
        "customer_id":"3f122513-757d-4ba0-b4dd-4f375b1344d9",
        "number":12345,
        "csv":123,
        "name_card":"lucas card",
        "validity":"08/31"
    }
    ```
``Update Customer registration`` 

To update for customers, we have several routes. The ID is the id field from the t_customer table
* Update all customer fields
    *  Url: http://localhost:5000/api/customer/ID?customer=true
    * Method: PUT
    * Body:
        ```json
        {
            "name":"teste",
            "email":"teste@gmail",
            "contact":121233,
            "cpf":123
        } 
        ```
* Update one customer passaword field
    * Url: http://localhost:5000/api/customer/ID?customer=true&pass=true
    * Method: PUT
    * Body:
        ```json
        {
            "password": "123"
        } 
        ```
To update to the customer address. The ID is the id field from the t_address table
* Update all address fields
    * Url: http://localhost:5000/api/customer/ID?address=true
    * Method: PUT
    * Body:
        ```json
        {
            "street":"teste",
            "block":"teste",
            "number":23,
            "state":"EUA"
        } 
        ```
To update to the customer card. The ID is the id field from the t_credit_card table
* Update all address fields
    * Url: http://localhost:5000/api/customer/ID?card=true
    * Method: PUT
    * Body:
        ```json
        {
            "number":12343,
            "csv":133,
            "name_card":"teste card",
            "validity":"08/34"
        } 
        ```
``Delete Customer``

To update for customers, we have several routes. The ID is the id field from the t_customer table
* Url: http://localhost:5000/api/customer/ID
* Method: DELETE
