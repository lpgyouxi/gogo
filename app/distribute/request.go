package distribute


import (
	"github.com/gin-gonic/gin"
)

/**
分配：链接上下文
link context
*/

// 定义一个接受参数的函数类型
type AppReqestFunc func(*AppReqest)

//定义参数结构体
type AppReqest struct {
	 *gin.Context
}

/* 
定义参数赋予方法  
*/
func Handle(r AppReqestFunc) gin.HandlerFunc {
	// gin.HandlerFunc 理解为一种可操作的函数接口，匿名函数
	return func(c *gin.Context) {
		r(&AppReqest{c})
	}
}

/* 
定义参数获取法  
*/

//定义获取get参数
func (params AppReqest) InputGet(paramName string) string {
	return params.DefaultQuery(paramName, "")
}


/* 
定义返回方法 
*/

//成功返回
func (c *AppReqest) Success(data interface{}) {
	c.JSON(Success(data))
}
