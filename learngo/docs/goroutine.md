
非抢占

# goroutine 定义

* 任何函数只需加上go就能送给调度器运行
* 不需要在定义时区分是否是异步函数
* 调度器在合适的点进行切换
* 使用 `-race` 来检测数据访问冲突

# goroutine 可能的切换点

* I/O, select
* channel
* 等待锁
* 函数调用（有时）
* runtime.Gosched()

只是参考，不能保证切换，不能保证在其他地方不切换

```shell
go run -race *.go
```