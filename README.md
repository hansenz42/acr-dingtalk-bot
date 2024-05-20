# 🤖 Go! 阿里云 ACR 镜像服务中转钉钉通知机器人服务

阿里云免费个人版的 ACR 镜像服务不提供钉钉机器人通知的功能。只有企业版提供。

使用本程序可以将免费版中镜像构建的触发器通知，转换为钉钉机器人消息，推送到钉钉。

![image](https://github.com/hansenz42/aliyun-acr-dingtalk-notifier/assets/11825586/a8620d43-0743-4986-94a2-a854c412295d)

## 准备

- 公网访问的云服务器/容器平台/serverless
- 开通阿里云 ACR 服务并创建一个镜像仓库
- 钉钉帐号

## 使用

### 1 设置 ACR 触发器

进入阿里云控制台 -> ACR -> 个人实例 -> (你创建好的镜像仓库) -> 触发器 -> 创建

![image](https://github.com/hansenz42/aliyun-acr-dingtalk-notifier/assets/11825586/87a321cd-cad7-42f9-b510-cdfd2baa8eec)

- 名称：随便起一个，例如 "success"
- URL：填写 `http(s)://{你的ip地址}:8085/OTFlNDNlOGItNzc0NS00ZTczLWFjMzYtMGEzYTI0MzExY2Vl`
- 触发方式：全部触发或按照你的需求修改

### 2 设置钉钉机器人

- 在手机端新建一个只有自己一个人的群（仅手机端可以）：
  右上角加号 -> 创建群聊 -> 普通群

- 创建一个自定义机器人并配置（仅电脑端可以，吐槽一下创建机器人的流程有点过于复杂）：
  打开群聊 -> 右上角齿轮图标 -> 机器人 -> 添加机器人 -> 自定义

- 机器人设置：
  安全设置中填写你的服务器的 IP 地址，然后复制 Webhook 地址

### 3 main.go 中修改地址

将代码中的 `TARGET_URL` 地址改刚刚复制的 webhook 地址

### 4 编译并运行项目

```bash
# 编译应用程序
go build

# 运行应用程序
./acr_to_dingtalk
```

检查一下服务器的防火墙配置。如果想折腾可以部署一个守护，或者用个反向代理

## 自定义访问地址

如需要服务的 http 接口地址，可以自行在此处修改

```go
// 修改改字符串修改接口地址，ACR 触发器 url 也要修改
r.POST("/OTFlNDNlOGItNzc0NS00ZTczLWFjMzYtMGEzYTI0MzExY2Vl", func(c *gin.Context) {
  ...
}
```