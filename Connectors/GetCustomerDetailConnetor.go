package Connectors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"serverLocalGo/Models"

	"github.com/gin-gonic/gin"
)

func GetCustomerDetail(c *gin.Context, cif string) (err error) {
	var reqDetails Models.RequestGetDetails

	reqDetails.CIF = cif
	reqDetails.TransactionID = "testgolangdaripostman"

	fmt.Println("reqDetails.CIF : ", reqDetails.CIF)

	c.BindJSON(&reqDetails)

	reqDetailsByte := new(bytes.Buffer)
	json.NewEncoder(reqDetailsByte).Encode(reqDetails)

	var apiUrl = "https://apidev.banksinarmas.com/internal/accounts/management/v2.0/getDetail"
	client := http.Client{}

	req, err := http.NewRequest("POST", apiUrl, reqDetailsByte)
	fmt.Println(err)
	fmt.Println(req)
	if err != nil {
		return
	}

	req.Header = http.Header{
		"Content-Type":     {"application/json"},
		"X-Gateway-APIKey": {"a50cfeba-8c04-4a84-aa84-eb0191019606"},
	}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server Internal Error")
		c.Abort()
	} else {
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("failed to read body: %s", err)
		}
		c.String(http.StatusOK, string(response))
	}
	return
}
