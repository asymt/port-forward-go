# port-forward-go

#### 介绍
一个用go语言开发的简单的端口转发小工具

#### 使用说明

1.  拉取代码
2.  使用`go build port-forward-go`编译项目
3.  使用port-forward-go help查看使用帮助,如下
```
port-forward cli

Usage:
  port-forward [flags]
  port-forward [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print the version number of port-forward CLI

Flags:
  -b, --bind-address string   bind address (default "0.0.0.0")
  -h, --help                  help for port-forward
  -p, --local-port int16      local port (default 9001)
  -r, --remote-host string    remote host
  -P, --remote-port int16     remote port

Use "port-forward [command] --help" for more information about a command.
```


#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

