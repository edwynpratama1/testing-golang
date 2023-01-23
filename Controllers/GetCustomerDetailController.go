package Controllers

import (
	"fmt"
	"serverLocalGo/Managers"

	"github.com/gin-gonic/gin"
)

func GetDetail(c *gin.Context) {

	err := Managers.CustomerDetail(c)
	if err != nil {
		fmt.Println("error controller")
		fmt.Println(err)
	}
}
