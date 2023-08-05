module github.com/JuanGQCadavid/now-project/services/locationDataUpdater

go 1.17

require (
	github.com/JuanGQCadavid/now-project/services/pkgs/common v0.0.0-00010101000000-000000000000
	github.com/aws/aws-lambda-go v1.28.0
	github.com/go-sql-driver/mysql v1.7.0
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
	gorm.io/gorm v1.25.2 // indirect
)

replace github.com/JuanGQCadavid/now-project/services/pkgs/common => ../pkgs/common
