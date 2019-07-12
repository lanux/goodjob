package cas

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/lanux/goodjob/v1/common/consts"
)

type Interceptor interface {
	PreAuthentication(ctx iris.Context) bool
	PostAuthentication(ctx iris.Context, u interface{})
	BeforeLogout(ctx iris.Context)
}

type DefaultInterceptor struct {
	S *sessions.Sessions
}

// 认证前执行的方法，返回false结束认证操作
// 返回true 跳过认证
func (c *DefaultInterceptor) PreAuthentication(ctx iris.Context) bool {
	session := c.S.Start(ctx)
	return session.Get(consts.USER_SESSION_KEY) != nil
}

// 认证后处理方法
func (c *DefaultInterceptor) PostAuthentication(ctx iris.Context, u interface{}) {
	if u != nil {
		session := c.S.Start(ctx)
		session.Set(consts.USER_SESSION_KEY, u)
	}
}

// 退出登录前处理方法
func (c *DefaultInterceptor) BeforeLogout(ctx iris.Context) {
	session := c.S.Start(ctx)
	session.Delete(consts.USER_SESSION_KEY)
	c.S.Destroy(ctx)
}
