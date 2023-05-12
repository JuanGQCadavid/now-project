module github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker

go 1.20

require (
	github.com/JuanGQCadavid/now-project/services/pkgs/common v0.0.0-00010101000000-000000000000
	github.com/aws/aws-lambda-go v1.40.0
	github.com/google/uuid v1.3.0
)

require (
	github.com/aws/aws-sdk-go v1.44.245 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace github.com/JuanGQCadavid/now-project/services/pkgs/common => ../pkgs/common
