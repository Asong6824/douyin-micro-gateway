package global

import(
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/app/client/discovery"
)

var (
	Registry registry.Registry
	Resolver discovery.Resolver
)