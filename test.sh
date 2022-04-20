go run main.go  -cpuprofile cpu.prof -memprofile mem.prof
go tool pprof -http=:9999 cpu.prof