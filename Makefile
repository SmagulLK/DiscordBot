.PHONY: build run

build:
	go build -o app cmd/main.go

run:
	./app
#clean docker-build docker-run docker-clean
##clean:
#	rm -f app
#
#docker-build:
#	docker build -t ds .
#
#docker-run:
#	docker run -d --name dsgo ds
#
#docker-clean:
#	docker stop dsgo
#	docker rm dsgo
