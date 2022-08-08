package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

//定义一个全局变量 R
var R *gin.Engine

//路由结构体
type Route struct {
	Method     reflect.Value  //方法路由
}
//定义全局的路由集合
var Routes =[]Route{}
//加载路由
func Load(r *gin.Engine) {
	R = r
	fmt.Println("执行3")
	Registers()
}

func Registers() {
	for _,route :=range Routes{
		if len(Routes) > 0 {
			arguments := make([]reflect.Value, 0)
			//arguments[0] = reflect.ValueOf(route.Method).Call(arguments) // *gin.Context
			//reflect.ValueOf(route.Method).Method(0).Call(arguments)
			route.Method.Call(arguments)
		}
	}
}

//路由注册()
func Register(routerRegister interface{}) {
	ctrlName := reflect.TypeOf(routerRegister).String()
	fmt.Println("ctrlName=", ctrlName)
	className := reflect.ValueOf(routerRegister)
	fmt.Println(className.NumMethod())
	for i:=0;i<className.NumMethod();i++ {
		method := className.Method(i)
		action := className.Type().Method(i).Name
		fmt.Println(action)
		fmt.Println(method)
		router:=Route{Method: method}
		Routes = append(Routes, router)
	}
}
