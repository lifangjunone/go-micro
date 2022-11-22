# proto usage
## proto code generate command
```shell

protoc -I=. --go_out=. --go_opt=module="github.com/lifangjunone/go-micro" --go-grpc_out=. --go-grpc_opt=module="github.com/lifangjunone/go-micro" apps/*/pb/*.proto


protoc -I=. -I=./common/pb   --go_out=. --go_opt=module="github.com/lifangjunone/go-micro"  --go-grpc_out=. --go-grpc_opt=module="github.com/lifangjunone/go-micro"  apps/*/pb/*proto  common/pb/*/*proto

```