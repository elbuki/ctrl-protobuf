# This should compile the protocol buffers for the supported languages.
gen-proto:
	rm -f proto/*.pb.go
	cd src/ && protoc --go_out=../proto/ --go-grpc_out=../proto/ *.proto