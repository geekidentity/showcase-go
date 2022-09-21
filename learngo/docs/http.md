# http服务器性能分析

* import _ "net/http/pprof"
* 访问 /debug/pprof
* 使用 og tool pprof分析性能

在net/http/pprof中可以查看帮助文档
// Then use the pprof tool to look at the heap profile:

//	go tool pprof http://localhost:6060/debug/pprof/heap

// Or to look at a 30-second CPU profile:

//	go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30