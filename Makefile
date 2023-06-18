gen-proto:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean:
	rm -rf ./src/hello/pb/
	rm -rf ./src/hashtable/pb/

hello-service:
	go run ./src/hello

build-hashtable-service:
	go build -o hashtable-service ./src/hashtable 

hashtable-service: gen-proto build-hashtable-service
	./hashtable-service 8080