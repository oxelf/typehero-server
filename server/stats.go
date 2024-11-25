package server

import "github.com/gin-gonic/gin"

func (s *Server) getStats(c *gin.Context) {
	stats, err := s.db.GetStats()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	stats.SiteViews = s.count

	c.JSON(200, stats)
	return
}
