# minmin-tiktok
极简版抖音后端 kitex+hertz

字节青训项目，但是我接口重复功能的接口就不写了
先这样吧后面我还会优化一下写一个v2的版本

完整版在 [mini_tiktok](https://github.com/hanyongyan/mini-tiktok)

### 技术栈
- 主语言：go
- orm框架：gorm + gorm.gen
- 数据库：mysql + redis
- http框架：hertz
- rpc框架：kitex
- 服务注册与服务发现：nacos
- idl：thrift
- 文件存储：腾讯cos
- 链路追踪：opentelemetry + jaeger
- 安装部署：docker compose
- 监控数据分析：Grafana + VictoriaMetrics

### 目录结构
```
.
├── README.md
├── cmd                 # 逻辑目录
│   ├── api             # 网关
│   ├── user            # 用户 
│   └── video           # 视频
├── docker-compose.yaml # 项目配置部署
├── go.mod
├── idl                 # 接口文件
│   ├── gateway
│   ├── user
│   └── video
├── kitex_gen           # kitex生成的文件
│   ├── userservice
│   └── videoservice
└── pkg
    ├── cache            # init
    ├── configs          # 数据库配置
    ├── consts           # 配置目录
    ├── dal              # gorm/gen 生成的文件  
    ├── errno            # 错误输出
    ├── mw               # 中间件 打印rpc信息
    └── utils            # jwt
```

### 架构
```
                                    http
                           ┌────────────────────────┐
 ┌─────────────────────────┤                        ├───────────────────────────────┐
 │                         │       apigateway       │                               │
 │      ┌──────────────────►                        │◄──────────────────────┐       │
 │      │                  └───────────▲────────────┘                       │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                           resolve                                 │       │
 │      │                              │                                    │       │
req    resp                            │                                   resp    req
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                   ┌──────────▼─────────┐                          │       │
 │      │                   │                    │                          │       │
 │      │       ┌───────────►        Nacos       ◄─────────────────┐        │       │
 │      │       │           │                    │                 │        │       │
 │      │       │           └────────────────────┘                 │        │       │
 │      │       │                                                  │        │       │
 │      │     register                                           register   │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
┌▼──────┴───────┴───┐                                           ┌──┴────────┴───────▼─┐
│                   │───────────────── req ────────────────────►│                     │
│        User       │                                           │        Video        │
│                   │◄──────────────── resp ────────────────────│                     │
└───────────────────┘                                           └─────────────────────┘
    thrift kitex                                                       thrift kitex
```


### Start
- User
  ```bash
    cd cmd/user
    sh build.sh
    sh output/bootstrap.sh  #这里生成的可能会多一个空行有的运行失败删掉即可
  ```
- Video
  ```bash
    cd cmd/video
    sh build.sh
    sh output/bootstrap.sh  #这里生成的可能会多一个空行有的运行失败删掉即可
  ```
- api
  ```bash
    cd cmd/api
    run go .
  ```


coding 过程在 [极简抖音笔记 - 夏末linxx的专栏](https://juejin.cn/column/7193694833058250811) 里