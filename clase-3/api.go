package main

import (
	"strconv"

	"./matem"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/calculate", func(c *gin.Context) {
		value := c.Query("value")
		v, err := strconv.Atoi(value)

		if err == nil {
			sum := matem.Sum(v)
			c.JSON(200, gin.H{
				"sum": sum,
			})
		} else {
			c.JSON(500, gin.H{
				"error": err,
			})
		}
	})

	r.Run()
}
