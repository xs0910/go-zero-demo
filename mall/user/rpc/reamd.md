运行命令，通过proto文件生产代码
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.