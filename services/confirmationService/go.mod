module github.com/JuanGQCadavid/now-project/services/confirmationService

go 1.20

require github.com/JuanGQCadavid/now-project/services/pkgs/common v0.0.0-00010101000000-000000000000

require github.com/neo4j/neo4j-go-driver/v4 v4.4.7 // indirect

replace github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder => ../pkgs/credentialsFinder

replace github.com/JuanGQCadavid/now-project/services/pkgs/common => ../pkgs/common
