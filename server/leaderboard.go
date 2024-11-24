package server

import (
	"fmt"
	"strconv"
	"typehero_server/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) getLeaderboard(c *gin.Context) {
    language, exists := c.GetQuery("language")
    if !exists {
        c.AbortWithStatus(400)
        return
    }

    mode, exists := c.GetQuery("mode")
    if !exists {
        c.AbortWithStatus(400)
        return
    }

    wordAmountString, exists := c.GetQuery("wordAmount")
    if !exists {
        c.AbortWithStatus(400)
        return
    }

    wordAmount, err := strconv.Atoi(wordAmountString)
    if err != nil {
        c.AbortWithStatus(400)
        return
    }

    page := 0
    pageString, exists := c.GetQuery("page")
    if exists {
        pageInt, err := strconv.Atoi(pageString)
        if err != nil {
            c.AbortWithStatus(400)
            return
        }
        page = pageInt
    }

    req := models.LeaderboardRequest{
        Language: language,
        WordAmount: wordAmount,
        Mode: mode,
        Page: page,
    }

    leaderboard, err := s.db.GetLeaderboard(req)
    if err != nil {
        fmt.Printf("error: %v \n", err)
        c.AbortWithStatus(500)
        return
    }

    c.JSON(200, leaderboard)
}
