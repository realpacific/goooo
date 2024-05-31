package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goooo/bloc"
	"goooo/config"
	. "goooo/logging"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var r = gin.Default()

const RANGE_REGEX = `^(\d+)(:\d+){0,1}$`

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
				"data": bloc.Factorial(number),
			})
		}
	})
	r.GET("/slice", func(c *gin.Context) {
		paramNumber, hasValue := c.GetQuery("range")
		if !hasValue {
			paramNumber = "0:0"
		}
		rangeRegex, _ := regexp.Compile(RANGE_REGEX)
		if !rangeRegex.MatchString(paramNumber) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Should be of format number:number but provided %v", paramNumber),
			})
		} else {
			split := strings.Split(paramNumber, ":")
			from, _ := strconv.Atoi(split[0])
			var to int
			if len(split) == 2 {
				to, _ = strconv.Atoi(split[1])
			} else {
				to = from
			}
			Logger.Infof("From %v, to %v", from, to)
			if from <= to {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("end range should higher than start range but is %v", paramNumber),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"data": bloc.Take(from, to),
				})
			}
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
