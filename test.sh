go run main.go  -cpuprofile cpu.prof -memprofile mem.prof
go tool pprof -http=:9999 cpu.prof

wget http://localhost:8080/debug/pprof/trace?seconds=5

wget http://localhost:8080/debug/pprof/profile