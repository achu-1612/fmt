build:
	GOOS=js GOARCH=wasm go build -o ui/main.wasm
run: docker
	docker run -it -p 3000:80 go-wasm
docker:
	docker build -t go-wasm .
