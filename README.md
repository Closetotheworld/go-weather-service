# Go-weather-service

외부의 간단한 응답 api를 이용한 summary를 생성하는 web server

## Getting Started
### Prerequisites

```
Go 1.18.x
```

### Installing

```
cd $GOPATH
mkdir -p src/Closetotheworld
git clone https://github.com/Closetotheworld/go-weather-service.git
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