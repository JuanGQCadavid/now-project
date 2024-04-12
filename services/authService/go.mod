module github.com/JuanGQCadavid/now-project/services/authService

go 1.20

require (
	github.com/JuanGQCadavid/now-project/services/pkgs/common v0.0.0-00010101000000-000000000000
	github.com/aws/aws-lambda-go v1.46.0
	github.com/aws/aws-sdk-go v1.51.21
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

replace github.com/JuanGQCadavid/now-project/services/pkgs/common => ../pkgs/common
