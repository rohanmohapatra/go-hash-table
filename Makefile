gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean:
	rm -rf ./src/hello/pb/

hello-service:
	go run ./src/hello

build-hashtable-service:
	go build ./src/hashtable

hashtable-service: build-hashtable-service
	./hashtable 8080