First install the next package

go get github.com/stretchr/testify

then:

* Create a file whose name would be <exisitngFileToTest>_test.go
* inside the file import 

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

* The methods on the file should be as

func Test[StructName| Func name][.FuncNameOFStruct]

* to run test we should:

go test ./... -coverprofile=coverage.out

go tool cover -func=coverage.out # to display coverage information on the terminal console.

go tool cover -html=coverage.out # to display the coverage information on HTML format in browser.
