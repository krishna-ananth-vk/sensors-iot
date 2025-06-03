build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o sensors-service main.go

run:
	go run main.go
