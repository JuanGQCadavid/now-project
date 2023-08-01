module github.com/JuanGQCadavid/now-project/services/locationDataUpdater

go 1.17

require (
	github.com/JuanGQCadavid/now-project/services/pkgs/common v0.0.0-00010101000000-000000000000
	github.com/aws/aws-lambda-go v1.28.0
	github.com/go-sql-driver/mysql v1.6.0
)

replace github.com/JuanGQCadavid/now-project/services/pkgs/common => ../pkgs/common
