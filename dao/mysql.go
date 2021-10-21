package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"test/models"
)

var (
	db *gorm.DB
)

// Initmysql 连接数据库
func Initmysql(mysqlconf *models.Mysqlconf)(err error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		//viper.GetString("mysql.user"),
		//viper.GetString("mysql.password"),
		//viper.GetString("mysql.host"),
		//viper.GetInt("mysql.port"),
		//viper.GetString("mysql.database"),

		mysqlconf.User,
		mysqlconf.Password,
		mysqlconf.Host,
		mysqlconf.Port,
		mysqlconf.Database,
	)

	db, err = gorm.Open("mysql", dsn)
	if err != nil{
		return errors.Wrap(err,"mysql init err")
	}


	//测试相应
	return db.DB().Ping()

}

// Close 对外暴露的关闭
func Close(){
	db.Close()
}

