package dao

import (
	"github.com/lanux/goodjob/v1/common/logger"
	"github.com/lanux/goodjob/v1/db"
)

type UserDao struct {
	BasicDao
}

type BasicDao struct {
}

func (*BasicDao) Create(m interface{}) bool {
	ok := db.Instance().NewRecord(m)
	logger.Errorf("create object %s", ok)
	return ok
}

func (*BasicDao) Update(u interface{}) bool {
	if err := db.Instance().Save(u).Error; err != nil {
		logger.Errorf("update by id Err:%s", err)
		return false
	}
	return true
}

func (*BasicDao) GetById(id int64, u interface{}) interface{} {
	if err := db.Instance().First(u, id).Error; err != nil {
		logger.Errorf("GetUserByIdErr:%s", err)
		return nil
	}
	return u
}

func (*BasicDao) Delete(u interface{}) bool {
	if err := db.Instance().Delete(u).Error; err != nil {
		logger.Errorf("Delete object error:%s", err)
		return false
	}
	return true
}
