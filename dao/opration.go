package dao

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"test/models"
)

func PutallDate() (err error,todolist []*models.Todo) {
	//这个定义的不是指针是一个存指针的切片
	if err = db.Find(&todolist).Error;err != nil {
		zap.L().Error("db.find all date err")
		return errors.Wrap(err,"find all date fail"),nil
	}
	return
}

func PutaDate(id string) (err error,todo *models.Todo) {
	//返回的是指针所以不能让指针指的地址发生改变
	todo = new(models.Todo)
	if err = db.Where("id=?",id).Find(todo).Error;err != nil{
		zap.L().Error("db.find a date err")
		return errors.Wrap(err,"find a date fail"),nil
	}
	return
}

func SaveDate(todo *models.Todo)(err error){
	err = db.Save(todo).Error
	zap.L().Error("db.save err")
	return errors.Wrap(err,"change status fail")
}

func CreateDate(todo *models.Todo)(err error){
	err = db.Create(todo).Error
	zap.L().Error("db.create err")
	return errors.Wrap(err,"create this date fail")
}

func DeleteDate(id string)(err error){
	err = db.Where("id=?",id).Delete(&models.Todo{}).Error
	zap.L().Error("db.delete err")
	return errors.Wrap(err,"delete this date fail")
}

