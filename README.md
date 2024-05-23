# MyShop REST API

This is a RESTful API for managing products in an online shop.

## Prerequisites

Before running this project, make sure you have the following installed:

- Go programming language (https://golang.org/)
- SQLite database (https://www.sqlite.org/index.html)

## Installation

1. Clone the repository:

```git clone https://github.com/cylegacy/myshop.git```

2. Navigate to the project directory:
   
```cd myshop```

3. Install dependencies:
   
```go mod tidy```


## Usage

To run the project, execute the following command:

```go run cmd/server/main.go```


The server will start and listen for incoming requests on port 8080 by default.

## Accessing Swagger Documentation

Once the server is running, you can access the Swagger documentation for the API by navigating to:

http://localhost:8080/swagger/index.html


This will open up the Swagger UI where you can explore and interact with the API endpoints.

## Endpoints

- GET /api/products: Retrieve a list of all products with optional filtering and pagination.
- GET /api/products/{id}: Retrieve a product by its ID.
- POST /api/products: Create a new product.
- PUT /api/products/{id}: Update an existing product.
- DELETE /api/products/{id}: Delete a product by its ID.

  
## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.


