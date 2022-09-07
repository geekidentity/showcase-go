# 单元测试

## 命令行


```shell
# 执行当前目录下所有单元测试
go test .
```

```shell
# 单元测试输出代码覆盖率
go test -coverprofile=c.out

go tool cover -html=c.out
```

```shell
# benchmark性能测试
go test -bench .  -cpuprofile cpu.out

# cpu.out 为二进制文件，查看方式
go tool pprof cpu.out
```