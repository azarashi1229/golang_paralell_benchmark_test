# golang_paralell_benchmark_test

以下でベンチマークを使って実行可能。
```
cd main
go test -bench .
```

以下、自分のローカルPCでの実行結果になります。実行環境はMacOS 13.3.1(a),Apple M1,16GB. 
```
% go test -bench .
goos: darwin
goarch: arm64
pkg: awesomeProject/main
BenchmarkMutex-8        87194834                13.58 ns/op
BenchmarkAtomic-8       173514421                6.886 ns/op
PASS
ok      awesomeProject/main     3.912s
```

GoRurutineを追加したパターン
```
go test -bench .
goos: darwin
goarch: arm64
pkg: awesomeProject/main
BenchmarkMutex-8         3736531               341.2 ns/op
BenchmarkAtomic-8        4187598               362.8 ns/op
PASS
ok      awesomeProject/main     3.631s

```
