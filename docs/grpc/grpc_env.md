## 1. 安装protobuf
IOS
```shell
brew info protobuf
brew install protobuf
```
win
下载地址：https://github.com/protocolbuffers/protobuf/releases/tag

### 验证
```shell
protoc --version
```
## 2. 安装gRPC
```shell
go get google.golang.org/grpc
```
## 3. 安装第三方依赖
```shell
go get -u google.golang.org/protobuf/proto
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# go get github.com/favadi/protoc-go-inject-tag
```
