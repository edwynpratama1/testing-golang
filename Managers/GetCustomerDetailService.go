package Managers

import (
	"fmt"
	"serverLocalGo/Configs"
	"serverLocalGo/Connectors"

	"github.com/gin-gonic/gin"
)

func CustomerDetail(c *gin.Context) (err error) {

	mobileNumber := c.Params.ByName("mobileNumber")
	fmt.Println("mobileNumber Service : ", mobileNumber)

	dataDB := Configs.DatabaseOracle(mobileNumber)
	if dataDB == "" {
		fmt.Println("Failed connect to Database")
		fmt.Println(err)
		return err
	} else {
		err = Connectors.GetCustomerDetail(c, dataDB)
		if err != nil {
			fmt.Println("Faied get data")
			fmt.Println(err)
			return err
		}
	}
	return
}
