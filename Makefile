# a not so complicated makefile :P

build:
	go build -o bin/ghm src/*.go

rpi:
	env GOOS=linux GOARCH=arm go build -o bin/ghm src/*.go
