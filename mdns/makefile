all: build

build: 
	cd server && make
	cd client && make

run:
	server/server &
	client/client 
