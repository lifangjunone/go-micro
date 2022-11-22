# go-micro
go micro service platform
## generate proto file command
```shell
protoc -I=. --go_out=. --go_opt=module="github.com/lifangjunone/go-micro" --go-grpc_out=. --go-grpc_opt=module="github.com/lifangjunone/go-micro" apps/*/pb/*.proto

```