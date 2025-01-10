build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

run: 
	go build
	./pwa-resume