# jwt+单点登录接口，服务端无状态存储
1. 这是一个用Golang编写的单点登录
2. 请务必在内网中运行
3. 添加了jwt的状态控制，使用内存的缓存和redis缓存
4. 添加了注册中心

### 快速开始 cd jwtService
1. go build main.go
2. main.exe # linux系统等使用: sh main
3. main.exe -help 可以查看提示