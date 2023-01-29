# Simple CMS Admin

Welcome to the Simple CMS Admin Service. An open-source Content Management System based on the echo framework. As admin, we can use the features provided by this service in the form of management of articles and categories. By using this service we can insert, update, delete and get details of each item (article and category). This service has implemented clean architecture principles, a practical software architecture solution from Robert C. Martin (known as Uncle Bob).

## Getting Started

### Prerequisites

- [Go 1.19.3](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

### Installation

- Clone the git repository:

```
$ git clone https://github.com/Assyatier21/simple-cms-admin.git
$ cd simple-cms-admin
```

- Install Dependencies

```
$ go mod tidy
```

- Create `config` folder in root path, then create a file `connection.go` in that folder containing this following code:

```
package config

const (
	User     = "YOUR_USERNAME_HERE"
	Password = "YOUR_PASSWORD_HERE"
	Host     = "localhost"
	Port     = "5432"
	Database = "YOUR_DATABASE_HERE"
	Schema   = "YOUR_SCHEMA_HERE"
	Sslmode  = "disable"
)
```

alternatively, we can just run this following command using makefile:

```
$ make all
```

### Running

```
$ go run cmd/main.go
```

### Features

This service has the following API endpoints:

- `admin/v1/articles`: get list of articles
- `admin/v1/article`: insert, update, delete and get details of article (method: POST, PATCH, DELETE, and GET)
- `admin/v1/categories`: get list of categories
- `admin/v1/category`: insert, update, delete and get details of category (method: POST, PATCH, DELETE, and GET)

We can test the endpoint using the postman collection in `simple-cms-admin/tools`.

### Testing

```
$ go test -v -coverprofile coverage.out ./...
```

## Install Local Sonarqube

please follow this [tutorial](https://techblost.com/how-to-setup-sonarqube-locally-on-mac/) as well.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/Assyatier21/simple-cms-admin/blob/master/LICENSE) file for details.
