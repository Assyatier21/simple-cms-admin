all: config

config:
	mkdir -p config
	echo "package config\n\nconst (\n\tUser     = \"YOUR_USERNAME_HERE\"\n\tPassword = \"YOUR_PASSWORD_HERE\"\n\tHost     = \"localhost\"\n\tPort     = \"5432\"\n\tDatabase = \"YOUR_DATABASE_HERE\"\n\tSchema   = \"YOUR_SCHEMA_HERE\"\n\tSslmode  = \"disable\"\n)" > config/connection.go


run: 
	go run cmd/main.go