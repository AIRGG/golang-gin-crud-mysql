package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"myapp/dbnya"
	"myapp/utils"
)

func GetElektronikBebasNamanya(c *gin.Context) {
	responseData := map[string]interface{}{
		"pesan": "",
		"error": true,
		"data":  nil,
	}
	rows, err := dbnya.DB.Query("SELECT * FROM tbl_buku")
	if err != nil {
		fmt.Println("ini eroor: " + err.Error())
		responseData["pesan"] = "[ERROR NIH B]: " + string(err.Error())
		c.JSON(http.StatusOK, responseData)
		return
	}
	defer rows.Close()
	data := utils.UtilParsingDB(rows, err)
	responseData["data"] = data
	responseData["error"] = false
	c.JSON(http.StatusOK, responseData)
}
