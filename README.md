# Simple CMS Admin

Welcome to the API documentation for the simple CMS Service. This API allows you to get an article and category as admin. This service using echo framework as well.

## Getting Started

### Prerequisites

- [Go 1.19.3](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

### Installation

- Clone the git repository:

```
git clone https://github.com/Assyatier21/simple-cms-admin.git
cd simple-cms-admin
```

- Install Dependencies

```
go mod tidy
```

- Create `config` folder in root path, then create a file `connection.go` in that folder with the following contents:

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

### Running

```
go run cmd/main.go
```

### API Endpoints Documentation

The API has the following endpoints:

- `/v1/articles`: get list of articles
- `/v1/article`: insert, update, delete and get details of article (method: POST, PATCH, DELETE, and GET)
- `/v1/categories`: get list of categories
- `/v1/category`: insert, update, delete and get details of category (method: POST, PATCH, DELETE, and GET)

We can test the endpoint using the collection located in : `simple-cms-admin/tools` (not updated yet).

### Testing

```
go test -v -coverprofile coverage.out ./...
```

## Install Local Sonarqube

please follow this [tutorial](https://techblost.com/how-to-setup-sonarqube-locally-on-mac/) as well.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/Assyatier21/simple-cms-admin/blob/master/LICENSE) file for details.
