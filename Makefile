build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o main main.go
	zip main.zip main
