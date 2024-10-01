# GoWit-BE-Case-Study

## Ticket Allocating and Purchasing Service

This project implements a REST API for allocating and purchasing tickets. The API ensures that ticket allocations do not drop below zero and that users cannot purchase more tickets than are available.

## API Endpoints

### 1. Create Tickets

-  **Endpoint:**  `POST /tickets`
-  **Request Body:**
```json
{
	"name": "Event Name",
	"description": "Event Description",
	"allocation": 100
}
```

-  **Response:**
- 201 Created
```json
{
	"id": 1,
	"name": "Event Name",
	"description": "Event Description",
	"allocation": 100
}
```
- 400 Bad Request if the input is invalid or cannot create ticket
```json
{
    "detail": "There was a problem processing your request, please try again",
    "status": 400,
    "title": "Bad Request"
}
```

### 2. Get Ticket by ID

-  **Endpoint:**  `GET /tickets/{id}`
-  **Response:**
- 200 OK on success
```json
{
	"id": 1,
	"name": "Event Name",
	"description": "Event Description",
	"allocation": 100
}
```
- 404 Not Found if the ticket does not exist
 ```json
{
    "detail": "Not found, please try again",
    "status": 404,
    "title": "Not Found"
}
```

### 3. Purchase Tickets

-  **Endpoint:**  `POST /tickets/{id}/purchases`
-  **Request Body:**
```json
{
	"quantity": 2,
	"user_id": "406c1d05-bbb2-4e94-b183-7d208c2692e1"
}
```
-  **Response:**
- 200 OK on success
- 400 Bad Request if the requested quantity exceeds the available allocation
```json
{
    "detail": "There was a problem processing your request, please try again",
    "status": 400,
    "title": "Bad Request"
}
```
## Technologies Used
- Go for API development
- PostgreSQL for data persistence
- GORM as the ORM
- Docker for containerization

## Local Development Setup

**Prerequisites**
- Go
- Docker
- PostgreSQL

**Running the Application**
To clone the repository:
`git clone https://github.com/zzeyneperdenn/GoWit-BE-Case-Study.git`

**Build and run the service using Docker Compose:**
`make up`

The API will be available at http://localhost:8080

**Testing**
To run the tests: `make test`
