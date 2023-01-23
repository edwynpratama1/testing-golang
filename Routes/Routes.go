package Routes

import (
	"net/http"
	"serverLocalGo/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp1 := r.Group("api/")
	{
		grp1.GET("welcome", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Hello localhost!")
		})
		grp1.POST("getdetail/:mobileNumber", Controllers.GetDetail)
	}

	return r
}
