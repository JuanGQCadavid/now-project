
## Golang commands

go get **package**

## To deploy on lambda

``` PowerShell
$env:GOOS = "linux"
$env:CGO_ENABLED = "0"
$env:GOARCH = "amd64"
go build -o main main.go
go build -o main .\cmd\lambda\main.go
```

* zip main and deployed it.
  