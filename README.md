# dmicro

## 目录
- [快速开始](#快速开始)
- [UMS统一消息系统](#UMS统一消息系统)
- [分布式事务处理](#分布式事务处理)
- [DRONE CI CD环境搭建指南](#DRONE CI CD环境搭建指南)
  - [vagrant环境搭建](docs/drone/vagrant.md)
  - [gitea安装配置](docs/drone/gitea.md)
  - [私有仓库harbor安装配置](docs/drone/harbor.md)
  - [drone安装配置](docs/drone/drone.md)

## 快速开始
### jaeger
```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:latest
```
### nats-streaming
```
docker run -d --name nats-streaming \
  -p 4222:4222 \
  -p 8222:8222 \
  -v /var/lib/docker-data/nats-streaming:/datastore \
  nats-streaming:latest -store file -dir datastore
```
### etcd
```
docker run -d --name etcd \
  -p 2379:2379 \
  -v /var/lib/docker-data/etcd:/etcd-data \
  quay.io/coreos/etcd:latest \
  etcd \
  --data-dir=/etcd-data \
  --name node1 \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379
```
### database
```
import sql scripts
```

### redis
```
docker run -d --name redis -p 6379:6379 redis
```

### gate micro web
```
cd gate/micro
go build
./micro web
```
### web
```
cd web/dd
go build
./dd
```
### gid
```
cd srv/gid
go build
./gid
```
### passport
```
cd srv/passport
go build
./passport
```
### user
```
cd srv/user
go build
./user
```
### postman
接口地址注意大小写，默认端口8082
```
curl -X POST \
  'http://localhost:8082/dd/passport/SmsLogin' \
  -H 'Content-Type: application/json' \
  -d '{
    "appid": 1,
    "plat": 1,
    "mobile": "13705918888",
    "code": "123456"
}'
```
以上使用的是web服务作为聚合api接口，还可以使用api服务作为聚合api接口

### gate micro api
```
./micro api --handler=api
```

### api
```
cd api/dd
go build
./dd
```

### postman
接口地址注意大小写，默认端口8080
```
curl -X POST \
  'http://localhost:8080/dd/passport/SmsLogin' \
  -H 'Content-Type: application/json' \
  -d '{
    "appid": 1,
    "plat": 1,
    "mobile": "13705918888",
    "code": "123456"
}'
```

## UMS统一消息系统
### 架构
<img src="https://github.com/fztcjjl/dmicro/raw/master/docs/ums.png">

### 运行micro
```
cd gate/micro
go build
./micro web
```

### 统一消息系统网关
```
cd web/ws
go build
./ws
```

### 统一消息系统逻辑服务器
```
cd srv/ums
go build
./ums
```

### 运行聊天应用服务器
```
cd srv/chat
go build
./chat
```

### 运行web chat demo
```
cd examples/chat
用浏览器打开chat.html文件
```

## 分布式事务处理

本方案基于本地消息表实现，示意图如下

<img src="https://github.com/fztcjjl/dmicro/raw/master/docs/capx.png">

## DRONE CI CD环境搭建指南
- [vagrant环境搭建](docs/drone/vagrant.md)
- [gitea安装配置](docs/drone/gitea.md)
- [私有仓库harbor安装配置](docs/drone/harbor.md)
- [drone安装配置](docs/drone/drone.md)

效果图：

<img src="https://github.com/fztcjjl/dmicro/raw/master/docs/drone/img/drone/7.png">