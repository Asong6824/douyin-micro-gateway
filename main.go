package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/Asong6824/douyin-micro-gateway/pkg/setting"
	"github.com/Asong6824/douyin-micro-gateway/global"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	//"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	//"github.com/Asong6824/douyin-micro-gateway/biz/rpc"
	"github.com/Asong6824/douyin-micro-gateway/biz/rpc"
	"github.com/minio/minio-go/v6"
	"io"
	"os"
)

func Init() {
	err := setupSetting()
	if err != nil {
		panic(err)
	}
	err = setupLogger()
	if err != nil {
		panic(err)
	}
	err = setupNacos()
	if err != nil {
		panic(err)
	}
	err = setupRpcClient()
	if err != nil {
		panic(err)
	}
}

func main() {
	Init()
	h := server.Default(
		server.WithStreamBody(global.EngineSetting.WithStreamBody),
		server.WithHostPorts(global.EngineSetting.WithHostPorts),
		server.WithRegistry(global.Registry, &registry.Info{
			ServiceName: global.EngineSetting.Registry.ServiceName,
			Addr:        utils.NewNetAddr("tcp", global.EngineSetting.Registry.Addr),
			Weight:      global.EngineSetting.Registry.Weight,
			Tags:        global.EngineSetting.Registry.Tags,
		}),
	)
	register(h)
	h.Spin()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Engine", &global.EngineSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Minio", &global.MinioSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    fileWriter := io.MultiWriter(f,os.Stdout)
    hlog.SetOutput(fileWriter)
	return nil
}

func setupNacos() error {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("172.17.0.2", 8848),
	}
	
	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
	}
	
	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return err
	}
	global.Registry = nacos.NewNacosRegistry(cli)
	global.Resolver = resolver.NewNacosResolver(cli)

	return nil
}

func setupRpcClient() error {
	err := rpc.SetupUserClient()
	if err != nil {
		return err
	}
	return nil
}

func setupMinioClient() error {
	var err error
	global.MinioClient, err = minio.New(global.MinioSetting.Endpoint, global.MinioSetting.AccessKeyID, global.MinioSetting.SecretAccessKey, global.MinioSetting.UseSSL)
    if err != nil {
        return err
    }
	return nil
}
