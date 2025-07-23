# Kratos 实例项目

本项目是经典分布式系统中的一网关+多个工作服的API网关模式设计（用户如果有需要也可以扩展网关）
在本系统中， 通过一个网关来统一管理和路由请求到多个后端服务，简化了客户端与服务之间的交互，同时提供了更高的灵活性和扩展性。

核心组件
API网关（gateway(app/gateway)）：
作为客户端和后端服务之间的中间层。
负责请求路由、协议转换、负载均衡、安全认证等功能。
网关服使用市面上流行的gin取代kratos默认的网关
网关服还集成了1个websocket功能

工作服务本项目中包含2个，包括一个user服和一个device服
可以通过网关服通过RPC调用到user服或者device服
user服和device服间也可以相互调用

客户端：
客户端只与网关交互，不直接访问后端服务。

工作服务中本人工gorm替代了kratos默认的orm（gorm市面上流行更广、性能强大、文档全）
<img width="755" height="341" alt="image" src="https://github.com/user-attachments/assets/e41e4125-9c28-49f7-9726-9b660993dc9f" />

关键词 kratos gin  gorm   websocket

如有问题可联系本人邮箱：zenghao306@163.com  

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

