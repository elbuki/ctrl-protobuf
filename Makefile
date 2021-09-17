# This should compile the protocol buffers for the supported languages.
proto:
	rm -f src/golang/*.pb.go
	protoc -I src/ src/*.proto --go_out=src/golang --go-grpc_out=src/golang