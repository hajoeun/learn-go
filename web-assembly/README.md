# Web Assembly with GO

## Compile

```bash
cd client
GOOS=js GOARCH=wasm go build -o main.wasm
```

## Run Server

```bash
go run main.go
# listening on ":8080"...
```

### Reference

- [https://github.com/golang/go/wiki/WebAssembly](https://github.com/golang/go/wiki/WebAssembly)
