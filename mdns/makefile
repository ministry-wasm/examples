all: build

build: 
	cd server && make
	cd client && make
	cd clienservert && make

run:
	server/server &
	client/client 
	cd clientserver && ./run8.sh
