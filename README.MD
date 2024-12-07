# GO CLEAN ARCH

## RUN SERVER APPLICATION
<h4>go to server directory and run</h4>

## LOAD PACKAGES
```bash
go mod tidy
```

```golang
go run main.go
# OR
nodemon --exec go run main.go --signal SIGTERM
```

## USING GO VERSION
``1.23.3``

## PROJECT STRUCTURE
### USING CLEAN ARCHITECTURE

    ├── assets
    │   └── assets
    ├── config
    │   └── env.go
    ├── constant
    │   ├── code_status.go
    │   └── constant.go
    ├── docker-compose.yaml
    ├── Dockerfile-development
    ├── go.mod
    ├── go.sum
    ├── INSTALLATION.MD
    ├── main.go
    ├── models
    │   ├── models.go
    │   └── paginator.go
    ├── README.MD
    ├── routes
    │   └── api.go
    ├── server
    │   └── server.go
    └── services
        └── service
            ├── handler
            │   ├── auth_handler.go
            │   └── auth_http.go
            ├── handler.go
            ├── mocks
            │   └── AuthRepository.go
            ├── README.MD
            ├── repository
            │   ├── auth_repository.go
            │   └── auth_repository_test.go
            ├── repository.go
            ├── usecase
            │   └── auth_usecase.go
            └── usecase.go

    13 directories, 25 files
