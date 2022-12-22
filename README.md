# Simple CMS User

Welcome to the API documentation for the simple CMS Service. This API allows you to get an article and category as user. This service using echo framework as well.

## Getting Started

### Prerequisites

- [Go 1.19.3](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

### Installation

- Clone the git repository:

```
git clone https://github.com/Assyatier21/simple-cms-user.git
cd simple-cms-user
```

- Install Dependencies

```
go mod tidy
```

### Running

```
go run cmd/main.go
```

### API Endpoints Documentation

The API has the following endpoints:

- `/v1/articles`: get list of articles
- `/v1/article`: get details of article by id
- `/v1/categories`: get list of categories
- `/v1/category`: get details of category by id

We can test the endpoint using the collection located in : `simple-cms-user/tools`.

### Testing

```
go test -v -coverprofile coverage.out ./...
```

## Install Local Sonarqube

please follow this [tutorial](https://techblost.com/how-to-setup-sonarqube-locally-on-mac/) as well.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/Assyatier21/simple-cms-user/blob/master/LICENSE) file for details.
