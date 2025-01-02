# Intern_Backend

This repository contains the backend code for the Intern project. It includes various functionalities and APIs to manage and interact with the database, such as adding, updating, deleting, and retrieving data. The project is built using Go and utilizes several libraries and frameworks, including Gin for the web framework, GORM for ORM, and Swagger for API documentation.

## Features

- Add, update, delete, and retrieve data from the database
- User authentication and authorization
- API documentation with Swagger
- Configuration management with environment variables

## Getting Started

To get started with this project, follow the instructions below:

### Prerequisites

- Go 1.19 or later
- MySQL database

### Installation

1. Clone the repository:
  ```sh
  git clone https://github.com/yourusername/Intern_Backend.git
  cd Intern_Backend
  ```

2. Install the dependencies:
  ```sh
  go mod tidy
  ```

3. Set up the environment variables in the `.env` file:
  ```env
  API_SECRET=your_api_secret
  TOKEN_HOUR_LIFESPAN=24
  DBUSER=root
  DBPASS=""
  DBHOST=localhost
  DBPORT=3306
  DBNAME=industrial
  ```

4. Run the application:
  ```sh
  go run main.go
  ```

### API Documentation

The API documentation is available at `http://localhost:8080/swagger/index.html` after running the application.

### API Endpoints

#### Authentication

- **Login Admin**
  - `POST /auth/admin/login`
  - Request Body: `{ "kode": "string" }`
  - Response: `{ "token": "string" }`

#### Product Management

- **Search Products**
  - `GET /get-product/search`
  - Query Params: `nama`
  - Response: `[{ "id": "int", "nama_barang": "string", "kategori_barang": "string" }]`

- **Filter Products by Category**
  - `GET /get-product/filter`
  - Query Params: `kategori`
  - Response: `[{ "id": "int", "nama_barang": "string", "kategori_barang": "string" }]`

- **Add Product**
  - `POST /update-product/add`
  - Request Body: `{ "nama_barang": "string", "kategori_barang": "string" }`
  - Response: `{ "message": "Product added successfully" }`

- **Delete Product**
  - `DELETE /update-product/delete`
  - Request Body: `{ "id": "int" }`
  - Response: `{ "message": "Product deleted successfully" }`

- **Update Product**
  - `PUT /update-product/update`
  - Request Body: `{ "id": "int", "nama_barang": "string", "kategori_barang": "string" }`
  - Response: `{ "message": "Product updated successfully" }`

## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request. We welcome all contributions!

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.