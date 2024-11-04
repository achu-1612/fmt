build:
	GOOS=js GOARCH=wasm go build -o ui/main.wasm
run: docker
	docker run -it -p 3000:80 go-wasm
docker:
	# goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
	docker build -t go-wasm .
