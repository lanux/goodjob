package auth

import (
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"github.com/lanux/goodjob/v1/common/consts"
	"github.com/lanux/goodjob/v1/config"
	"github.com/lanux/goodjob/v1/db"
	"github.com/lanux/goodjob/v1/web/middleware/cas"
	"net/http"
)

func New(ssessions *sessions.Sessions) *Casbin {
	enforcer := casbin.NewEnforcer(config.CasbinFilePath, &Adapter{db.Instance()})
	enforcer.EnableLog(false)
	return &Casbin{enforcer: enforcer, s: ssessions}
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

func (c *Casbin) Check(ctx context.Context) bool {
	session := c.s.Start(ctx)
	userInfo := session.Get(consts.USER_SESSION_KEY)
	if userInfo == nil {
		return false
	}
	user := userInfo.(cas.AuthSuccessStruct)
	method := ctx.Method()
	path := ctx.RequestPath(false)
	return c.enforcer.Enforce(user.User, path, method)
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
	model["g"]["g"].Policy = append(model["g"]["g"].Policy, []string{"zhenlong.zhong", "root"})
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
