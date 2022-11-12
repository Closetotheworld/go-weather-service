# Go-weather-service

외부의 간단한 응답 api를 이용한 summary를 생성하는 web server

## Getting Started
### Prerequisites

```
Go 1.18.x
```

### Run (at your local stage)

```
// PORT option is optional. default value is 8080 

cd $GOPATH
mkdir -p src/Closetotheworld
git clone https://github.com/Closetotheworld/go-weather-service.git
go run main.go -k {API_KEY} -p {PORT}
```

### Run (at your docker stage)
```
// PORT option is optional. default value is 8080

docker build . -t weather-api   
docker run --rm -p {PORT_WHAT_YOU_WANT}:{PORT} --name weather-api weather-api -k {API_KEY} -p {PORT}
```

## Running the tests

### formatting

```
go fmt
```

### testing
```
go test -v ./...
```

## Built With

* [Gin-gonic](https://github.com/gin-gonic/gin/) - The web framework used
* [Cobra-cli](https://github.com/spf13/cobra) - to run server with flag

## Authors

* **Wonryang heo** - [Closetotheworld](https://github.com/Closetotheworld)