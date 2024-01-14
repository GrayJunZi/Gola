SHELL_PATH = /bin/sh
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/sh,/bin/sh)

# 1. 清理依赖包
tidy:
	go mod tidy
#	go mod vendor

# 2. 运行本地所有服务
run-local:
	go run app/services/sales-api/main.go
	