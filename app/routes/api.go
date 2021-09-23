package routes



import (
    "github.com/gin-gonic/gin"
    "gogo/app/http"       //业务模块
    "gogo/app/distribute"
)

func Route(e *gin.Engine) {

  	//test
  	e.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "hello",
        })
    })
    // 业务 
    test := http.TestController{}
    e.GET("/test", distribute.Handle(test.Test))
    e.GET("/hello", distribute.Handle(test.Hello))
}

