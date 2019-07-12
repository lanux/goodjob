package auth

import (
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"github.com/lanux/goodjob/v1/db"
	"github.com/lanux/goodjob/v1/web/middleware/cas"
	"net/http"
)

func InitCasbin(app *iris.Application, ssessions *sessions.Sessions) {
	enforcer := casbin.NewEnforcer("./config/casbinmodel.conf", &Adapter{db.Instance()})
	enforcer.EnableLog(true)
	casbinMiddleware := &Casbin{enforcer: enforcer, s: ssessions}
	app.Use(casbinMiddleware.ServeHTTP)
	//app.WrapRouter(casbinMiddleware.Wrapper())
}

type Casbin struct {
	enforcer *casbin.Enforcer
	s        *sessions.Sessions
}

func (c *Casbin) ServeHTTP(ctx context.Context) {
	if !c.Check(ctx) {
		ctx.StatusCode(http.StatusForbidden) // Status Forbiden
		ctx.StopExecution()
		return
	}
	ctx.Next()
}

//func (c *Casbin) Wrapper() router.WrapperFunc {
//	return func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
//		if !c.Check() {
//			w.WriteHeader(http.StatusForbidden)
//			w.Write([]byte("403 Forbidden"))
//			return
//		}
//		router(w, r)
//	}
//}

func (c *Casbin) Check(ctx context.Context) bool {
	session := c.s.Start(ctx)
	user := session.Get("user").(cas.AuthSuccessStruct)
	method := ctx.Method()
	path := ctx.RequestPath(false)
	return c.enforcer.Enforce(user.Attributes.ACCOUNT, path, method)
}

type Adapter struct {
	Db *gorm.DB
}

type CasbinRule struct {
	db.Model
	PolicyType string `gorm:"size:100"`
	Subject    string `gorm:"size:100"`
	Object     string `gorm:"size:100"`
	Action     string `gorm:"size:100"`
}

// 设置CasbinRule的表名为`casbin_rules`
func (CasbinRule) TableName() string {
	return "casbin_rules"
}

func (ad *Adapter) LoadPolicy(model model.Model) error {
	sec := "p"
	var rules []CasbinRule
	ad.Db.Find(&rules)
	for _, rule := range rules {
		model[sec][rule.PolicyType].Policy = append(model[sec][rule.PolicyType].Policy, []string{rule.Subject, rule.Object, rule.Action})
	}
	return nil
}

func (*Adapter) SavePolicy(model model.Model) error {
	panic("implement me")
}

func (*Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	panic("implement me")
}

func (*Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	panic("implement me")
}

func (*Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	panic("implement me")
}
