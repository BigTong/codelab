# rpc bentchmark

## background
- 
## benchmark result
- go test -bench="GrpcUpdateUser" -benchtime=10s 
- BenchmarkGrpcUpdateUser-8           100000            133371 ns/op

- go test -bench="JsonrpcUpdateUser" -benchtime=10s
- BenchmarkJsonrpcUpdateUser-8              50000            341992 ns/op

- go test -bench="JsonSerialAndDeSelrial" -benchtime=10s
- BenchmarkJsonSerialAndDeSelrial-8           200000            101162 ns/op

- go test -bench="ProtoSerialAndDeSelrial" -benchtime=10s
- BenchmarkProtoSerialAndDeSelrial-8        1000000             10166 ns/op

