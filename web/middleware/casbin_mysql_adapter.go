package cas

import (
	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
)

type Adapter struct{}

func (*Adapter) LoadPolicy(model model.Model) error {
	persist.LoadPolicyLine("", model)
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
