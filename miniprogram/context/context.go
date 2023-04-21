package context

import (
	"github.com/lucidapp/wechatsdk/v2/credential"
	"github.com/lucidapp/wechatsdk/v2/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
