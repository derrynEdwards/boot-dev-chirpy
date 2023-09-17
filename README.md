Chirpy
==============

Overview
--------
Go guided project from [Boot.dev](https://boot.dev)

Chirpy is a social network similar to twitter. In this project, we're building
a fully-fledged web server from scratch using the net/http standard library and 
extending it with Chi.

Learning Goals
--------------
- Architecture of decoupling back-end and front-end
- RESTful APIs
- Dynamic parameters in both query and url.
- HTTP Requests and Responses

Execution
---------
1. Clone this repository to your local machine:
```shell
git clone <repository-url>
```

2. Navigate to the project directory:
```shell
cd boot-dev-chirpy
```

3. Compile the program:
```shell
go build
```

4. Run the program:
```shell
./boot-dev-chirpy
```

File Tree
---------
```shell
boot-dev-chirpy
├── README.md
├── assets
│   └── logo.png
├── go.mod
├── go.sum
├── handler_chirps_create.go
├── handler_chirps_delete.go
├── handler_chirps_get.go
├── handler_login.go
├── handler_refresh.go
├── handler_users_create.go
├── handler_users_get.go
├── handler_users_update.go
├── handler_webhooks.go
├── headers.go
├── index.html
├── internal
│   ├── auth
│   │   └── auth.go
│   └── database
│       ├── chirps.go
│       ├── database.go
│       ├── revocation.go
│       └── users.go
├── json.go
├── logger.go
├── main.go
├── metrics.go
├── readiness.go
└── reset.go

```
