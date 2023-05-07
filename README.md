# basic-lambda
getting started with Lambda and Go in AWS


## create the most basic Lambda using AWS example code
1. `go mod init github.com/robertocamp/basic-lambda`
2. develop code in main.go
3. Preparing a binary to deploy to AWS Lambda requires that it is compiled for Linux and placed into a .zip file:
  + build an executable called "bootstrap" at the root of the project
  + `GOOS=linux GOARCH=amd64 go build -o bootstrap main.go`
  + `zip lambda-handler.zip bootstrap`
## Links
https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html
https://github.com/aws/aws-lambda-go