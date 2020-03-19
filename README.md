# Quick Accounts through Protoc

By using a protobuf definition I can quickly spin up reliable RPCs, REST gateways and swagger files through protoc.

I have a MongoDB instance which the RPC connects to and is able to serve up the swagger at the endpoint:
http://localhost:8080/swaggerui/dist

The example uses cobra to wrap the commands to spin up the RPC and REST along with a docker-compose file for the MONGO instance.

To run all the services:
  - MONGO: docker-compose up --build
  - RPC:   go run main.go rpc -q=true
  - REST:  go run main.go rest
