package server

import (
	"fmt"
	"time"
	"typehero_server/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) postResults(c *gin.Context) {
	var result models.Result
	err := c.BindJSON(&result)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	realIp := c.GetHeader("x-real-ip")
	cfConnectingIP := c.GetHeader("CF-Connecting-IP")

	if result.WPM > 250 {
		c.AbortWithStatusJSON(403, gin.H{
			"error": "you really think you can type that fast? I don't think so",
		})
		fmt.Printf("fake attempt %d wpm, real ip: %s %s \n", result.WPM, realIp, cfConnectingIP)
		return
	}

	score, err := verifyCaptcha(result.CaptchaToken)
	fmt.Printf("score: %.2f \n", score)
	if err != nil || score < 0.5 {
		c.AbortWithStatusJSON(403, gin.H{
			"error": "nice try bogdan",
		})
		return
	}

	if len(result.UserName) > 16 || len(result.UserName) < 3 {
		c.AbortWithStatus(403)
		return
	}

	result.Date = time.Now()

	err = s.db.CreateResult(result)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.Status(200)
}
