package router

import (
	"fmt"
	"gin_cli/controller"
)

type UserRouter struct {
}

func init() {
	//注册路由
	fmt.Println("执行")
	Register(&UserRouter{})
}
func (u *UserRouter) Index()  {
	R.GET("/user",controller.UserController{}.Index)
}
