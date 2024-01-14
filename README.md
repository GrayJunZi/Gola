# Gola

---
ArdanLabs - Ultimate Go Web Services with Kubernetes

## 一、介绍

### Linux中安装Go

(1). 进入[官网下载](https://go.dev/dl)对应系统版本的Go语言SDK，例如目前最新版 `go1.21.6.linux-amd64.tar.gz`。

(2). 将文件解压缩。
```bash
tar -zxvf go1.21.6.linux-amd64.tar.gz
```

(3). 将文件夹移动到 `/usr/local/` 目录下。
```bash
mv go /usr/local/go
```

(4). 修改 `/etc/profile` 文件，追加如下命令：
```
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

(5). 查看Go版本，检查是安装成功。
```bash
go version
```

(6). 修改Go镜像源
```bash
go env -w Go111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 二、模块

### 1. 初始化项目

(1). 初始化Go模块。
```bash
go mod init github.com/grayjunzi/gola
```

(2). 清理Go依赖。
```bash
go mod tidy
```

### 2. 部署优先心态(deploy first mentality)

- **结构化日志**，将日志信息转化为人类可读的信息。使用UBer开发的`Zap`日志库。通过`argo`查看日志内容。

- **日志器**，不要为日志器使用任何类型的单例或全局变量。我们要做的是在main中构造logger，并将logger传递到代码中需要记录的地方。（我们会知道所有的东西都是在哪里构造的，我们可以遵循执行的路径以及如何共享事物建造的地方，将有助于以后的测试。）

- **命名约定**，我们在创建一个包的时候，应该非常清楚这个包提供了什么，每个包都是一个API，其中包括静态库、程序不同部分的防火墙(firewall)。当编写包的时候，会有依赖的问题，例如包含：`utils`、`helpers`等。如果在go中有包含`modules`、`types`等名称的包将注定是要失败的。不能在项目中创建**通用类型包**。

> **不要让事情变得容易做，而是让事情变得容易理解。**