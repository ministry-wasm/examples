all: build

build:
	go build . 
	ls -l

debug: build
	dlv debug github.com/ministry-wasm/examples/mdns/clientserver

