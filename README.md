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