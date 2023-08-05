module github.com/JuanGQCadavid/now-project/services/locationDataUpdater

go 1.17

require (
	github.com/JuanGQCadavid/now-project/services/pkgs/common v0.0.0-00010101000000-000000000000
	github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder v0.0.0-20230805234016-df5442ed42e3
	github.com/aws/aws-lambda-go v1.28.0
	github.com/go-sql-driver/mysql v1.7.0
	gorm.io/driver/mysql v1.5.1
	gorm.io/gorm v1.25.2
)

require (
	github.com/aws/aws-sdk-go v1.44.245 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/neo4j/neo4j-go-driver/v4 v4.4.4 // indirect
)

replace github.com/JuanGQCadavid/now-project/services/pkgs/common => ../pkgs/common

//replace github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFindes => ../pkgs/credentialsFinder

replace github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFindes => ../pkgs/credentialsFinder
