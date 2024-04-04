# Car Catalog API

This project implements a RESTful API for managing a catalog of cars. It allows users to perform CRUD operations (Create, Read, Update, Delete) on car records, as well as retrieve a list of cars with pagination support.

## Endpoints:
```
GET /cars: Retrieve a list of all cars with optional pagination parameters (offset and limit).
GET /cars/{id}: Retrieve details of a specific car by its ID.
POST /cars/create: Create a new car record.
POST /cars/update/{id}: Update details of a specific car by its ID.
DELETE /cars/delete/{id}: Delete a car record by its ID.
```

## Request and Response Formats:

For GET /cars and GET /cars/{id} endpoints, the response will be in JSON format containing details of one or more cars.
For POST /cars/create and POST /cars/update/{id} endpoints, the request body should contain JSON data representing the car details to be created or updated.
For DELETE /cars/delete/{id} endpoint, the car ID to be deleted is specified in the URL path.
Pagination:

The GET /cars endpoint supports pagination with optional offset and limit parameters.
The offset parameter specifies the number of items to skip before returning results.
The limit parameter specifies the maximum number of items to return.

## Swagger Documentation:

The API is documented using OpenAPI (Swagger) specification. The Swagger UI can be accessed at /swagger-ui endpoint when the server is running.

## Usage:
```
Clone the repository: git clone <repository-url>
Navigate to the project directory: cd CarCatalog
Build the project: go build
Run the executable: ./CarCatalog
Access the API using the provided endpoints.
```

## Dependencies:
```
Go 1.15+
Gorilla Mux (github.com/gorilla/mux)
PostgreSQL database
Environment Variables:

Ensure the following environment variables are set:

PORT: Port number for the HTTP server (e.g., 8080).
DB_CONNECTION_STRING: Connection string for the PostgreSQL database.
```