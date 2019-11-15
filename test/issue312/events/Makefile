regenerate:
	go install github.com/tron-us/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/gogo/protobuf/protobuf/ \
	--gogo_out=. events.proto 
