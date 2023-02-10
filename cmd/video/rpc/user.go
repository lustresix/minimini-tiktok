package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"mini-min-tiktok/kitex_gen/userservice/userservice"
	"mini-min-tiktok/pkg/consts"
	"mini-min-tiktok/pkg/mw"
)

var userService userservice.Client

func GetUserClient() userservice.Client {
	return userService
}

func initUser() {
	// 接入OpenTelemetry 链路追踪
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.VideoServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)

	// 服务端配置
	// 创建clientConfig的另一种方式
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(consts.NacosAddr, consts.NacosPort),
	}

	// 客户端配置
	cc := constant.ClientConfig{
		NamespaceId:         "public",           // ACM的命名空间Id
		TimeoutMs:           5000,               // 请求Nacos服务端的超时时间，默认是10000ms
		NotLoadCacheAtStart: true,               // 在启动的时候不读取缓存在CacheDir的service信息
		LogDir:              "/tmp/nacos/log",   // 日志存储路径
		CacheDir:            "/tmp/nacos/cache", // 缓存service信息的目录，默认是当前运行目录
		LogLevel:            "info",             // 日志默认级别，值必须是：debug,info,warn,error，默认值是info
		Username:            "nacos",            // Nacos服务端的API鉴权Username
		Password:            "nacos",            // Nacos服务端的API鉴权Password
	}

	// 创建服务发现客户端
	r, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	// 为IDL中定义的服务创建客户端
	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(resolver.NewNacosResolver(r)), // 解析器
		client.WithMuxConnection(1),                       // 最大连接数
		client.WithMiddleware(mw.CommonMiddleware),        // 打印信息
		client.WithInstanceMW(mw.ClientMiddleware),        // 服务端地址和超时信息
		client.WithSuite(tracing.NewClientSuite()),        // 选项套件
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}), // 为 rpcInfo 提供初始信息
	)

	if err != nil {
		panic(err)
	}
	userService = c

}
