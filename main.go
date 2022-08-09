package main

import (
	"fmt"
	"gin_cli/bean"
	"gin_cli/common"
	_ "gin_cli/docs"
	"gin_cli/middleware"
	"gin_cli/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

// @title 搭建gin脚手架
// @version 1.0 版本dsdd
// @description 描述:更方便的对接接口sds
// @BasePath /api/v1 基础路径
// @query.collection.format multi
func main() {
	initConfig()
	initDb()
	r := gin.Default()
	banner()
	//使用全局中间件
	r.Use(middleware.CorsMiddleware()) //全局路由
	router.Load(r)
	r.Run(viper.GetString("application.port"))
}

//初始化配置文件
func initConfig() {
	workDir, _ := os.Getwd() //获取当前路径
	fmt.Println(workDir)
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	fmt.Println(viper.GetString("application.port"))
	if err != nil {
		panic(err)
	}
}

func initDb() {
	// 用户名:密码@tcp(ip:port)/数据库?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var user = bean.User{Id: 1}
	db.Take(&user)
	fmt.Printf("%#v\n", user)
}

func banner() {
	if !(common.Isfile("./banner.txt")) {
		fmt.Println("当前文件不存在")
	} else {
		content, err := ioutil.ReadFile("./banner.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(content))
	}
}
