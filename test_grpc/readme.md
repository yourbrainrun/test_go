###### mac 安装
> brew install protoc-gen-go

###### 执行生成定义结构声明文件
```
protoc --go_out=./ --go-grpc_out=./ ./proto/*.proto

protoc --go_out=./ ./proto/*.proto
protoc --go_out=plugins=grpc:./ ./proto/*.proto
```
