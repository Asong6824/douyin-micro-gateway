package app

import (
	"github.com/Asong6824/douyin-micro-gateway/pkg/errno"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/app"
)

type Response struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func SendFailResponse(c *app.RequestContext, errno errno.ErrNo) {
	c.JSON(consts.StatusOK, Response{
		Code: errno.ErrCode,
		Msg:  errno.ErrMsg,
	})
	return
}