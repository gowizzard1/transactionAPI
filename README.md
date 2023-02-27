## TransactionAPI
This is a RESTful API for performing transactions and retrieving balance information.

### Dependencies
This project uses Golang v1.16, GORM v1.21, and PostgreSQL v13. To install these dependencies, refer to their respective installation guides.

### Installation
Clone the repository
```bash
git clone https://github.com/username/transactionAPI.git
```
Navigate to the project directory
```bash
cd transactionAPI
```
Create a .env file and populate it with the following environment variables:
```makefile
HOST=<db host>
DB_USER=<db username>
DB_NAME=<db name>
DB_PASSWORD=<db password>
DB_PORT=<db port>
```

Build the docker image
```
docker-compose build
```
Start the docker container
```
docker-compose up
```
Access the API at http://localhost:8080
Endpoints
Create Transaction
```bash
POST /create
```
Creates a new transaction.

Request Body
```json
{
"amount": 100,
"type": "credit"
}
```
Response
```json
{
"id": 1,
"amount": 100,
"type": "credit",
"balance": 100
}
```
Get Balance
```bash
GET /balance
```
Retrieves the current balance.

Response
```json
{
"balance": 100
}
```

