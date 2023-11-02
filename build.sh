protoc -I=proto \
   --go_out=proto --go_opt=paths=source_relative \
   --go-grpc_out=proto --go-grpc_opt=paths=source_relative \
   --grpc-gateway_out=proto --grpc-gateway_opt=paths=source_relative \
   helloworld/hello.proto


# client test
curl -X POST -k http://localhost:8891/v1/example/echo -d '{"name": " hello3"}'
curl -X POST -k http://localhost:8891/v1/example/test -d '{"number": 100}'