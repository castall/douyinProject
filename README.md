#### 目前默认端口

api：8200

etcd：2379

### 项目生成

当前api已生成完毕无需重复

~~在./cmd/api/ 下运行以下命令使用hertz命令生成api项目~~

```shell
cd douyinProject/cmd/api
go mod init douyinProject/api
hz new -idl ../../idl/api.thrift -module douyinProject/api
go mod tidy
```

在各自目录下根据./idl下的.thrift文件，使用kitex命令生成rpc项目

例如在./cmd/user/ 下使用以下命令生成rpc项目

```shell
cd douyinProject/cmd/user
kitex --module douyinProject/user --service user ../../idl/userrpc.thrift 
go mod tidy
```

### 项目编写

#### api service编写

每个人需要修改自己的 ./cmd/api/biz/handler/api/ 下的service内容

需要调用远程方法则需要参照以下规则编写自己的远程方法微服务

#### 微服务远程方法编写

1. 在./cmd下建立自己的目录
2. 在./idl下编写自己的.thrift文件
3. 使用kitex命令利用.thrift文件生成自己的rpc server
4. 使用kitex命令将自己的远程方法也添加到./cmd/api/项目中



写好自己的微服务远程方法后需要到./api目录下生成kitex-gen路径

下面以user服务的例子
```shell
# 注意需要填写-module指定使用api这个module
cd douyinProject/cmd/api
kitex -module douyinProject/api ../../idl/userrpc.thrift
# 拉取依赖
go mod tidy
```

### 项目启动

运行启动etcd注册中心

```shell
etcd -listen-client-urls="http://0.0.0.0:2379" --advertise-client-urls="http://0.0.0.0:2379"
```
启动各自微服务注册到注册中心

启动api，接受http请求并可通过etcd注册中心获取微服务远程调用方法，根据远程调用反馈结果返回用户请求

测试阶段可以直接通过api返回结果以供测试
