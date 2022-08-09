# Go-Template
### Technique using:
- Backend: Golang language (https://go.dev/doc/)
- Database: PostgreSql
- Manage database connect with GORM (https://gorm.io/docs/)
- Manage api router with Gin (https://gin-gonic.com/docs/)
### IDE using:
- Coding: GoLand (https://www.jetbrains.com/go/)
- View database: TablePlus (https://tableplus.com/)
- Use docker for start mysql container (https://docs.docker.com/)
- Test request api: Postman (https://www.postman.com/)
### Startup server:
- Open GoLand IDE and add configuration
- Open docker and start postgres container
- Go configs/template.yaml change portgres config
- Press run button or run `go run cmd/main.go` at terminal
- Finally, make request with http://127.0.0.1:8080/auth/login method POST & Header Content-Type=application/json & Body `{"username":"admin", "password": "123456"}`