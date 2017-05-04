# rpc bentchmark

## background
- 
## benchmark result
- go test -bench="." -benchtime=10s                                                    
- BenchmarkGrpcUpdateUser-8                100000            133138 ns/op
- BenchmarkJsonrpcUpdateUser-8               50000            351460 ns/op
- BenchmarkJsonSerialAndDeSelrial-8         200000             99101 ns/op
- BenchmarkFJsonSerialAndDeSelrial-8        200000             99367 ns/op
- BenchmarkProtoSerialAndDeSelrial-8       1000000             10393 ns/op