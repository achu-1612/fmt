build:
	GOOS=js GOARCH=wasm go build -o ui/main.wasm
run:
	goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'