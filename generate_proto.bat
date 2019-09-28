protoc -I api/ api/contract.proto --go_out=plugins=grpc:api
protoc -I api/ api/reservation.proto --go_out=plugins=grpc:api
protoc -I api/ api/saga.proto --go_out=plugins=grpc:api