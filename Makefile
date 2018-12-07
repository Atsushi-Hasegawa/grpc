
all: setup

setup:
	go get "github.com/golang/glog"
	go get "golang.org/x/net/context"
	go get "github.com/grpc-ecosystem/grpc-gateway/runtime"
	go get "google.golang.org/grpc"
	go get "github.com/gogo/protobuf/proto"
