package distribute


import (
	"github.com/gin-gonic/gin"
	"net/http"  
	"reflect"  // 对象操作的类
)

/**
成功返回，你可以建其它返回，如失败、错误
Successful return, you can create other returns, such as failure and error
*/
func Success(data interface{}) (int, gin.H) {
	ref := reflect.ValueOf(data)
	if ref.IsNil() {
		return http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    make(map[string]interface{}),
			"message": "请求成功",
		}
	} else {
		return http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    data,
			"message": "请求成功",
		}
	}
}



