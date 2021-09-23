package main
import (
	"github.com/gin-gonic/gin"
	// "net/http"
	"gogo/app/boot"
)


func main()  {
	server := gin.Default()

	

	boot.Bootstrap(server)

	// server.GET("/", func(c *gin.Context) {
	// 	param := c.DefaultQuery("name", "Guest")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "hello 测试通过",
	// 		"data":param,
	// 	})
	// })

	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}