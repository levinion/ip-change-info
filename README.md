# ip-change-info

## 功能

用以提醒用户公网 ip 变动，作为那些没有域名或没有配置 ddns 服务的用户连接内网的临时解决办法。

## 使用方式

1. 将文件下载到本地
```
git clone https://github.com/levinion/ip-change-info.git
```
2. 修改 main.go 文件中的相关配置
3. 运行 `go build`
4. 运行二进制程序
```
./ipChangeInfo
# 或后台挂起运行
nohup ./ipChangeInfo &
```