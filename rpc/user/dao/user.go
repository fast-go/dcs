package dao

import (
	"dcs/gen/gorm/model"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (d *UserDao) Create(m model.User) {
	d.db.Create(m)
}

//docker login -u service-1677292445749 -p 01da42e74f322b105e8bdb41eb13a2368396e6a2 dongweitiao-docker.pkg.coding.net
//
//dongweitiao-docker.pkg.coding.net/demo/service/demo:master-362ba8ada81a1b9e82c3f00102aca69aae3a1369

func (d *UserDao) FindOneById(id int64) (m model.User) {
	d.db.First(&m, id)
	//d.db.Debug().First(&m)
	return m
}
