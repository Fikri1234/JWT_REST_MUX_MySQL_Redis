# JWT_REST_MySQL_JWT_Redis
Web service CRUD using Golang with gorilla-MUX for create REST api, MySQL as database, Viper as environment variable, JWT for secure service, and redis to store token.


**Prerequisites**

1. [Go](https://golang.org/)
2. [Gorilla Mux](https://github.com/gorilla/mux)
3. [Mysql](https://www.mysql.com/downloads/)
4. [Viper](https://github.com/spf13/viper)
5. [SQLMock](https://github.com/DATA-DOG/go-sqlmock)
6. [Assert](https://godoc.org/github.com/stretchr/testify/assert)
7. [BCrypt](https://godoc.org/golang.org/x/crypto/bcrypt)
8. [JWT](https://github.com/dgrijalva/jwt-go)
9. [UUID](https://github.com/segmentio/ksuid)
10. [Redis](https://github.com/gomodule/redigo)

**Getting Started**
1. Firstly, we need to get MUX, MySQL, Viper, sqlmock, assert library dependencies and install it
```
go get github.com/gorilla/mux  
go get github.com/go-sql-driver/mysql
go get github.com/spf13/viper
go get github.com/DATA-DOG/go-sqlmock
go get github.com/stretchr/testify/assert
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/segmentio/ksuid
go get github.com/gomodule/redigo/redis
```
2. Import dump.sql to your MySQL and configure your credential in folder resource
![Alt text](asset/configureCredentialDB.PNG?raw=true "Configure your credential DB")
3. Open cmd in your project directory and type `go test -v` , you should get a response similar to the following:
![Alt text](asset/testing.PNG?raw=true "Response Unit Testing")

4. To run application,open cmd in your project directory and type
```
go run main.go
```
5. Download [Redis for Windows](https://github.com/dmajkic/redis/downloads)
6. After you download Redis, you’ll need to extract the executables and then double-click on the redis-server executable.

**Sample Payload**
1. [Login](asset/login.PNG)
2. [Logout](asset/logout.PNG)
3. [Get User By Id](asset/getUserById.PNG)
4. [Get User Detail By Id](asset/getUserDetailById.PNG)
5. [Get All User](asset/getAllUser.PNG)
6. [Get All User Detail](asset/getAllUserDetail.PNG)
7. [Create User](asset/createUser.PNG)
8. [Create User Detail](asset/createUserDetail.PNG)
9. [Update User](asset/updateUser.PNG)
10. [Update User Detail](asset/updateUserDetail.PNG)
11. [Delete User By Id](asset/deleteUserById.PNG)
12. [Delete User Detail By Id](asset/deleteUserDetailById.PNG)
13. [Example error response,in case Update User Detail](asset/updateUserDetailError.PNG)