package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xlzd/gotp"
)

type Otp struct {
	Otp string `json:"otp"`
}

func main() {
	r := gin.Default()

	totp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")

	r.GET("/:otp", func(c *gin.Context) {
		otp := c.Param("otp")
		timestamp := time.Now().UTC().Unix()
		result := totp.Verify(otp, timestamp)
		c.JSON(http.StatusOK, gin.H{
			"result":    result,
			"timestamp": timestamp,
		})
	})

	r.POST("/", func(c *gin.Context) {
		timestamp := time.Now().UTC().Unix()
		genOtp := totp.At(timestamp) // default expiration time 30s
		c.JSON(http.StatusOK, gin.H{
			"otp":       genOtp,
			"timestamp": timestamp,
		})
	})

	r.Run()
}
