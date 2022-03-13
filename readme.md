# Simple REST API Go App

## Features:
- Gin for the web framework
- Gorm for ORM
- jwt-go for authentication
- Docker with MySQL and phpmyadmin

## How to use:
- Clone the repo
- `sudo docker-compose -p local-prep --file docker-compose.yml up -d`
- Run build.sh
- `go run server.go`