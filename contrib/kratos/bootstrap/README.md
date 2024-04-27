# bootstrap

引导应用启动的相关设置。

## 配置

微服务或者说云原生应用的配置最佳实践是将配置文件和应用代码分开管理——不将配置文件放入代码仓库，也不打包进容器镜像，而是在服务运行时，把配置文件挂载进去或者直接从配置中心加载。Kratos的config组件就是用来帮助应用从各种配置源加载配置。

- 本地文件file
- 配置中心的适配：Consul、Etcd、Nacos、Apollo、Polaris、Kubernetes

[更多信息](https://go-kratos.dev/docs/component/config)

其他功能：

- 支持日志记录器：Stdout、Zap、Logrus、Fluent、Aliyun、Tencent
- 支持HTTP服务：Kratos Rest
- 支持GRPC服务端和客户端
- 支持数据库：Redis
- 支持服务注册发现中心：Consul、Etcd、Nacos、Eureka、Zookeeper、Kubernetes、Servicecomb
- 支持对象存储：Minio
- 支持链路追踪：Zipkin、otlp-http、otlp-grpc、Jaeger

