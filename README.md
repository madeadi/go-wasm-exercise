# Exercise on Go Wasm

```
GOOS=js GOARCH=wasm go build -o static/hash.wasm cmd/hash/main.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static
go run ./cmd/webserver/main.go 
```

Visit `localhost:3000/index.html`