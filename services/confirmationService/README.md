## Install common packages

go mod edit --replace github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder=../pkgs/credentialsFinder
go mod edit --replace github.com/JuanGQCadavid/now-project/services/pkgs/common=../pkgs/common

go mod tidy

## Install packages that could be usefull

go get "github.com/neo4j/neo4j-go-driver/v4/neo4j"  