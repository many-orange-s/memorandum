package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"test/models"
)

var Conf = new(models.Config)

func Init() (err error) {
	//设定文件的位置
	viper.SetConfigFile("./config.yml")
	//viper.SetConfigType("yml")
	//viper.AddConfigPath(".")

	//读取文件
	if err = viper.ReadInConfig();err != nil {
		if _,ok := err.(viper.ConfigFileNotFoundError);!ok {
			//文件找到了 但是产生了别的错误
			return errors.Wrap(err,"filefind but other err")
		}
	}

	if err = viper.Unmarshal(Conf);err != nil {
		return errors.Wrap(err,"setting unmarshal err")
	}

	fmt.Println(Conf.Host)
	//监视文件是否改动
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err = viper.Unmarshal(Conf)
		// 配置文件发生变更之后会调用的回调函数
		zap.L().Info("config.yml change")
		log.Println("配置文件发生了改变")
	})
	return
}
