# go_otel_error_case

2 个问题，

1. 无法从源码编译出 v0.8.0 的 otel 工具了。
复现步骤：
```
cd ~
// clone 源码
git clone https://github.com/alibaba/opentelemetry-go-auto-instrumentation.git
cd opentelemetry-go-auto-instrumentation
git fetch
git checkout v0.8.0
make build

// 检查 工具版本号
~/opentelemetry-go-auto-instrumentation/otel version
otel version 0.9.0_877aae4
```
编译出的产物，已经变成 0.9.0 版本。

2. 当无法从源码构建出 v0.8.0 otel 后，尝试使用 github release 页的 v0.8.0 工具编译项目，但是发生报错，之前是不会的。
复现步骤：
```
cd ~
git clone https://github.com/wangquanchn/go_otel_error_case.git
cd go_otel_error_case
// 下载 otel v0.8.0 工具
wget https://github.com/alibaba/opentelemetry-go-auto-instrumentation/releases/download/v0.8.0/otel-darwin-amd64
// 赋予可执行权限
chmod -R +x otel-darwin-amd64
// 编译报错
./otel-darwin-amd64 go build
``` 
编译时将会报错，报错信息如下：
```
➜  go_otel_error_case git:(main) ✗ ./otel-darwin-amd64 go build
===== Environments =====
Command    : ./otel-darwin-amd64 go build
ErrorLog   : .otel-build/preprocess/debug.log
WorkDir    : /Users/wangquan/go_otel_error_case
Toolchain  : darwin/amd64, go1.23.6, 0.8.0_877aae4c
===== Fatal Error ======
Error      : Failed to run command
Reason     : WORK=/var/folders/m9/s6dlmb9559x9cnzh7zxtbyk40000gn/T/go-build925988538
go.opentelemetry.io/otel/trace: /Users/wangquan/go_otel_error_case/otel-darwin-amd64: exit status 1

Cause      : goroutine 1 [running]:
runtime/debug.Stack()
	/Users/y1yang/.gvm/pkgsets/go1.22/global/pkg/mod/golang.org/toolchain@v0.0.1-go1.23.6.darwin-arm64/src/runtime/debug/stack.go:26 +0x5e
github.com/alibaba/opentelemetry-go-auto-instrumentation/tool/errc.New(0x3f6, {0xc000116000, 0xac})
	/Users/y1yang/Desktop/opentelemetry-go-auto-instrumentation/tool/errc/errcode.go:99 +0xc5
github.com/alibaba/opentelemetry-go-auto-instrumentation/tool/preprocess.runCmdCombinedOutput({0x0, 0x0}, {0xc001a24b40?, 0xc001812080?, 0x8bee500?})
	/Users/y1yang/Desktop/opentelemetry-go-auto-instrumentation/tool/preprocess/preprocess.go:172 +0xcb
github.com/alibaba/opentelemetry-go-auto-instrumentation/tool/preprocess.runBuildWithToolexec({0xc0000be1c0, 0x2, 0x2})
	/Users/y1yang/Desktop/opentelemetry-go-auto-instrumentation/tool/preprocess/preprocess.go:762 +0x555
github.com/alibaba/opentelemetry-go-auto-instrumentation/tool/preprocess.Preprocess()
	/Users/y1yang/Desktop/opentelemetry-go-auto-instrumentation/tool/preprocess/preprocess.go:999 +0x327
main.main()
	/Users/y1yang/Desktop/opentelemetry-go-auto-instrumentation/tool/cmd/main.go:181 +0x9b

Detail.command: [build -toolexec=/Users/wangquan/go_otel_error_case/otel-darwin-amd64 remix -work -a]
➜  go_otel_error_case git:(main) ✗
```
