// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(
		server.WithStreamBody(true),
		server.WithHostPorts("0.0.0.0:8000"),
	)

	register(h)
	h.Spin()
}
