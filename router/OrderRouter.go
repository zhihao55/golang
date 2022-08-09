package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
/**
	路由
 */
type OrderRouter struct {
}

func init() {
	//注册路由
	fmt.Println("执行")
	Register(&OrderRouter{})
}
func (o *OrderRouter) Index()  {
	R.GET("/tt", func(c *gin.Context) {
		c.JSON(http.StatusOK,"ssss")
	})
}