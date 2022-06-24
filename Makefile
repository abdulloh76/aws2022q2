.PHONY: build clean deploy

build:
	dep ensure -v
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/getProductsList getProductsList/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/getProductById getProductById/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/createProduct createProduct/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
