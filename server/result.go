package server

import (
	"time"
	"typehero_server/models"

	"github.com/gin-gonic/gin"
)


func (s *Server) postResults(c *gin.Context) {
    var result models.Result
    err :=  c.BindJSON(&result)
    if err != nil {
        c.AbortWithStatus(400)
        return
    }

    score, err := verifyCaptcha(result.CaptchaToken)
	if err != nil || score < 0.5 {
        c.AbortWithStatusJSON(403, gin.H{
        "error": "nice try bogdan",
        })
		return
	}

    if (len(result.UserName) > 16 || len(result.UserName) < 3) {
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
