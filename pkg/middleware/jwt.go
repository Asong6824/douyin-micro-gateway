package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/Asong6824/douyin-micro-gateway/pkg/errno"
	"github.com/Asong6824/douyin-micro-gateway/pkg/utils"
	pkgapp "github.com/Asong6824/douyin-micro-gateway/pkg/app"
)

func AuthRequiredToken() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token, ok := c.GetQuery("token")

		if !ok {
			token, ok = c.GetPostForm("token")
		}

		if !ok {
			hlog.CtxInfof(ctx, "token is not exist clientIP: %v\n", c.ClientIP())
			pkgapp.SendFailResponse(c, errno.AuthorizationFailedErr)
			c.Abort()
			return
		}
		hlog.CtxInfof(ctx, "token: %v clientIP: %v\n", token, c.ClientIP())

		claims, err := utils.CheckToken(token)

		if err != nil {
			hlog.CtxInfof(ctx, "token is invalid clientIP: %v\n", c.ClientIP())
			pkgapp.SendFailResponse(c, errno.AuthorizationFailedErr)
			c.Abort()
			return
		}
		c.Set("tokenid", claims.UserId)
		c.Next(ctx)
	}
}  

func AuthOptionalToken() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token, ok := c.GetQuery("token")
		if !ok {
			token, ok = c.GetPostForm("token")
		}
		if ok {
			hlog.CtxInfof(ctx, "token: %v clientIP: %v\n", token, c.ClientIP())

			claims, err := utils.CheckToken(token)
			if err != nil {
				hlog.CtxInfof(ctx, "token is invalid clientIP: %v\n", c.ClientIP())
				pkgapp.SendFailResponse(c, errno.AuthorizationFailedErr)
				c.Abort()
				return
			}

			c.Set("tokenid", claims.UserId)
		} else {
			hlog.CtxInfof(ctx, "token is not provided clientIP: %v\n", c.ClientIP())
		}

		c.Next(ctx)
	}
}