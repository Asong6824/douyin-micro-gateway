package rpc

import (

	"github.com/Asong6824/douyin-micro-gateway/kitex_gen/user/userservice"
	"github.com/Asong6824/douyin-micro-gateway/kitex_gen/user"
	"github.com/Asong6824/douyin-micro-gateway/global"
	"github.com/Asong6824/douyin-micro-gateway/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"context"
)

func SetupUserClient() error {
	c, err := userservice.NewClient(
		"user-service",
		client.WithResolver(global.Resolver),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: global.EngineSetting.Registry.ServiceName}),
	)
	if err != nil {
		return err
	}
	global.UserClient = c

	return nil
}

func Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	resp, err := global.UserClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}