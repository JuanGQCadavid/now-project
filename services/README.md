# Steps
## To start a project:

1. Create a folder
2. go mod init github.com/JuanGQCadavid/now-project/services/**ServiceName**

## To install a package

1. go get **package**

## Gin in lambda

### First, install the Lambda go libraries.
go get github.com/aws/aws-lambda-go/events
go get github.com/aws/aws-lambda-go/lambda

### Next, install the core library.
go get github.com/awslabs/aws-lambda-go-api-proxy/...