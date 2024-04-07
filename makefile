generate-server:
	@protoc   server/api/api.proto --go-grpc_out=server/pkg --go_out=server/pkg  --experimental_allow_proto3_optional
generate-client:
	@protoc   server/api/api.proto --go-grpc_out=controller/pkg --go_out=controller/pkg  --experimental_allow_proto3_optional