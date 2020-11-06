p=$(shell pwd)

test.unit.service:
	PROJ_DIR=$p go test -count=1 -mod=vendor -v ./pkg/service

test.all:
	test.unit.service

server:
	PROJ_DIR=$p GOFLAGS=-mod=vendor go run ./main.go server

heroku.push:
	heroku container:push web -a dcard-test --recursive

heroku.release:
	heroku container:release web -a dcard-test

heroku.deploy:
	make heroku.push
	make heroku.release
