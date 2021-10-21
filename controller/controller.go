package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test/dao"
	"test/models"
)

func IndexHander(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ShowallDate(c *gin.Context) {
	err,todolist := dao.PutallDate()
	if err != nil{
		log.Printf("%#v",err)
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		//这个有用返回表里面的所有信息
		c.JSON(http.StatusOK, todolist)
	}
}

// UpdateaDate 对后端来说没有什么操作的 其实就是把statue的状态改成1
func UpdateaDate(c *gin.Context) {
	//获取要传过来的数据id
	id , ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"err":"没有这个id"})
		return
	}

	//找到id的位置
	err, todo := dao.PutaDate(id)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{"err":err.Error()})
		return
	}

	//其实更新对于后端来说就是把status的状态改成1
	//他会把所有的信息全都传过来（前面传的id只是让你找那个数据的位置）只是把status改成了1
	//你只要重新绑定然后再一次储存就好了
	//之前返回的就是一个指针
	c.BindJSON(todo)

	//在原来的id的位置上进行改变
	//原数据的改变
	//如果用create就重新有一个数据
	err = dao.SaveDate(todo)
	if err != nil {
		log.Printf("%#v",err)
		c.JSON(http.StatusOK,gin.H{"err":err.Error()})
	}else{
		c.JSON(http.StatusOK,todo)
	}
}

func GetDate(c *gin.Context)  {
	var todo models.Todo
	//数据获取
	c.ShouldBind(&todo)

	err :=dao.CreateDate(&todo)
	if  err != nil {
		log.Printf("%#v",err)
		c.JSON(http.StatusOK, gin.H{"err ": err.Error()})
	} else {
		//做一个返回 其实没有什么用 可以不做
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteaDate(c *gin.Context) {
	id , ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"err":"没有这个id"})
		return
	}

	err := dao.DeleteDate(id)
	if err != nil {
		log.Printf("%#v",err)
		c.JSON(http.StatusOK,gin.H{"err":err.Error()})
	}else{
		c.JSON(http.StatusOK,gin.H{"status":"ok"})
	}
}
