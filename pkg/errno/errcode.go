package errno

import (
	
)

var (
	ServiceErr             = NewErrNo(10001, "Service is unable to start successfully")
	ParamErr               = NewErrNo(10002, "Wrong Parameter has been given")
	AuthorizationFailedErr = NewErrNo(10003, "Authorization failed")
)
