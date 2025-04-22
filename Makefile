build:
	GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
	zip function.zip bootstrap
deploy-stage:
	aws lambda update-function-code --function-name gin-api2 --zip-file fileb://function.zip
