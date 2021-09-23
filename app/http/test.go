package http

import (
	"github.com/gin-gonic/gin"
	"gogo/app/distribute"


	"fmt"
)

type TestController  struct {

}

func (r *TestController) Hello(request *distribute.AppReqest) {
	fmt.Println("Ok");
	request.Success(make([]int, 0));
}


func (r *TestController) Test(request *distribute.AppReqest) {
		request.JSON(200, gin.H{
			"message": "hello 测试通过",
		})
}

func (r *TestController) TestGet(request *distribute.AppReqest) {
		var data = []string{}
		var name = request.InputGet("name")
		data = append(data,name)
		request.Success(data);
	
}