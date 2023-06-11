# API Server  V2



Go-Gin后端框架重写Melo CMDB API Server


# 项目结构
```shell
├── api
│   └── v2
├── cmd
├── conf
├── core
├── dto
├── initialize
├── job
├── message
├── model
├── pojo
├── router
│   └── v2
├── scripts
├── service
└── util
```

| 文件夹 | 说明                    | 描述                        |
|----| ----------------------- | --------------------------- |
| `api` |api层|api层|
| `v2` |v2版本接口|v2版本接口|
| `cmd` |命令包|存放命令行相关函数|
| `conf` |配置包|存放相关配置文件|
| `core` |核心文件|核心组件（server）的初始化|
| `dto` |数据传输对象|存放对应的传输对象结构体|
| `initialize` |初始化|router的初始化|
| `job` |工作包|存放notification job、task job相应函数|
| `message` |消息包|job消息的传输，状态更改|
| `model` |模型层|模型对应数据表和CRUD方法|
| `pojo` |简单对象|存放model层与service层相应的传输对象的结构体和函数|
| `router` |路由层|路由层|
| `--v2` |v2版本路由|v2版本路由|
| `scripts` |脚本包|存放相关脚本|
| `service` |服务层|存放业务逻辑问题|
| `util` |工具包|工具函数封装|

# LICENSE

- API Server V2: [MIT license](https://mit-license.org/)