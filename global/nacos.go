package global

import(
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/kitex/pkg/discovery"
)

var (
	Registry registry.Registry
	Resolver discovery.Resolver
)