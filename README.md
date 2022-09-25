# Soccer API (Backend) - Golang
A soccer API

**Development by:** 
- Findryankp

## Start Project
- go run .

## Init Project
- go mod init soccerApi

## Install Packages
* go get github.com/labstack/echo/v4
* go get github.com/joho/godotenv

**Swaggo documentation API:** 
* go get -u github.com/swaggo/swag/cmd/swag
* go get -u github.com/swaggo/files
* go get -u github.com/swaggo/echo-swagger

## Step By step (MAC OS)
- export PATH=$(go env GOPATH)/bin:$PATH
- swag init

# Problem 2 Enhanced

- Type: Open problem

You want to publish a soccer API that can give a list of soccer players in a team. Also, don’t forget to create functionality to input the team and it’s player.

## Expected Feature

- Create a team
- Add player to the team
- Get a team, including with the players
- Get a player

## Terms and conditions

- Don't use any RDBMS or noSQLDB (limitation accepted: data can be lost when app / server is restarted)
- The API must be configurable (using config file)
- Solve this problem using language you know and confident
- Create unit test for the function
- If you can make it the API response faster, it will be a plus
- API response in json format
- We want you to make this API as easy as possible to start. No hassle when we try to running your code
- Don’t forget to make documentation

## Bonus point

- Implementation in Go.