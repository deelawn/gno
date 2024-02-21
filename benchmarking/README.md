
To enable benchmarking

```sh
export ENABLE_BENCHMARKING=true 
```

To print out opcode details on console
```sh
export OPCODE_DETAILS=true

./build/gnoland start | grep "benchmark.Op"
```
The benchmarking contract is located at `examples/gno.land/r/x/benchmark/ops.gno`

To execute the benchmarking functions, add and edit file located at `gno.land/genesis/genesis_txs.txt`
