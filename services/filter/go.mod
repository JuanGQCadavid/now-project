module github.com/JuanGQCadavid/now-project/services/filter

go 1.18

require (
	github.com/JuanGQCadavid/now-project/services/pkgs/common v0.0.0-00010101000000-000000000000
	github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder v0.0.0-00010101000000-000000000000
	github.com/aws/aws-lambda-go v1.28.0
	github.com/aws/aws-sdk-go v1.44.245
	github.com/awslabs/aws-lambda-go-api-proxy v0.12.0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-sql-driver/mysql v1.7.0
	github.com/google/uuid v1.3.0
)

require (
	git.sr.ht/~sbinet/gg v0.4.1 // indirect
	github.com/ajstarks/svgo v0.0.0-20211024235047-1546f124cd8b // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-fonts/liberation v0.3.1 // indirect
	github.com/go-latex/latex v0.0.0-20230307184459-12ec69307ad9 // indirect
	github.com/go-pdf/fpdf v0.8.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/neo4j/neo4j-go-driver/v4 v4.4.4 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/image v0.7.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	gonum.org/v1/plot v0.13.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
	gorm.io/gorm v1.25.2 // indirect
)

replace github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder => ../pkgs/credentialsFinder/

replace github.com/JuanGQCadavid/now-project/services/pkgs/common => ../pkgs/common
