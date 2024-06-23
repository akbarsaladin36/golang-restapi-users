<h1 align="center">Golang - REST API Users</h1>

This API is created by me as implementation from tutorial that i learned a week ago. You can use this project for free [More about Golang](https://go.dev/)

## Built With

[![Go/Golang](https://img.shields.io/badge/Golang-1.20-cyan.svg?style=rounded-square)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-v.1.10-blue.svg?style=rounded-square)](https://gin-gonic.com/docs/)
[![GORM](https://img.shields.io/badge/Gorm-v.1.25-blue.svg?style=rounded-square)](https://gorm.io/)


## Requirements

1. <a href="https://go.dev/">Go</a>
2. <a href="https://gin-gonic.com/docs/">Gin</a>
3. <a href="https://gorm.io/">GORM</a>
4. <a href="https://www.mysql.com/">MySQL</a>

## How to run the app ?

1. Clone this project
2. Open app's directory in CMD or Terminal
3. Turn on Web Server and MySQL can using Third-party tool like xampp, etc.
4. Create a database with the name #nama_database, and Import file sql to **phpmyadmin**
5. Open Postman desktop application or Chrome web app extension that has installed before
6. Choose HTTP Method and enter request url.(ex. localhost:3600/)
7. You can see all the end point [here](https://documenter.getpostman.com/view/14780095/2sA3XWdefr)
8. Type `go run main.go` to activated the server.

## Set up project

1. Open "setup.go" file in folder "database" on your favorite code editor, and you can change a port below:

```
Router.Run(":3600") // you can change that ":3600" with the port that you want
```

2. You can change username, password, and database name in file "setup.go" like code below:

```
dbc := "<username_database>:<password_database>@tcp(127.0.0.1:3306)/<database_name>?charset=utf8mb4&parseTime=True&loc=Local"
```

## Feature

1. Get All Data for Users
2. Get One Data for User
3. Create, Update, and Delete User


## License

© [Muhammad Akbar Saladin Siregar](https://github.com/akbarsaladin36/)
