# product-api

## Table of Content

- [How To Run This Project](#how-to-run-this-project)
- [Endpoints](#endpoints)
- [Tools Used](#tools-used)

### How To Run This Project

> - Make Sure you have run the db_product_api.sql in your mysql
> - Edit environment variables in the .env file

#### Run the Applications

```bash
# Move to directory
$ cd workspace

# Clone into your workspace
$ git clone https://github.com/rizkakhairani/product-api.git

# Move to project
$ cd product-api

# Run the application
$ go run main.go
```

### Endpoints

| Method | Endpoint                             | Description               |
| :----- | :----------------------------------- | :----------------------   |
| GET    | /products                            | Get list of all products  |
| GET    | /products/:id                        | Get product by id         |
| POST   | /products                            | Add product               |
| PUT    | /products/:id                        | Edit product by id        |
| DELETE   | /products/:id                        | Delete product by id      |
| ---    | ---                                  | ---                       |
| GET    | /products?sort_by=desc(created_at)   | Get list of all products sort by latest product    |
| GET    | /products?sort_by=asc(price)         | Get list of all products sort by lowest price      |
| GET    | /products?sort_by=desc(price)        | Get list of all products sort by highest price     |
| GET    | /products?sort_by=asc(name)          | Get list of all products sort by name (A-Z)        |
| GET    | /products?sort_by=desc(name)         | Get list of all products sort by name (Z-A)        |

### Tools Used

In this project, I use some tools listed below. All libraries listed in [`go.mod`](https://github.com/rizkakhairani/product-api/blob/main/go.mod).

- [Echo](https://echo.labstack.com/)
- [GORM](https://gorm.io/)
