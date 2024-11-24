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

    result.Date = time.Now()

    err = s.db.CreateResult(result)
    if err != nil {
        c.AbortWithStatus(500)
        return
    }

    c.Status(200)
}
