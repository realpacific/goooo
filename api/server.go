package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goooo/config"
	"goooo/services"
	"log"
	"net/http"
	"strconv"
)

var r = gin.Default()

func init() {
	r.GET("/fact", func(c *gin.Context) {
		paramNumber, hasValue := c.GetQuery("number")
		if !hasValue {
			paramNumber = "1"
		}
		number, err := strconv.Atoi(paramNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Invalid input %v", number),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": services.Factorial(number),
			})
		}
	})
}

func Serve() {
	port := config.GetPort()
	portInt := strconv.Itoa(port)
	err := r.Run(":" + portInt)

	if err != nil {
		log.Fatalln("failed to start server", err)
	}
}
