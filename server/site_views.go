package server

import (
	"fmt"
	"typehero_server/models"

	"github.com/gin-gonic/gin"
)


func (s *Server) increaseSiteViews(c *gin.Context) {
	defer s.countMutex.Unlock()

	s.countMutex.Lock()
	s.count++

	c.Status(200)
}

func (s *Server) saveSiteViews() {
	var siteViews models.SiteViews
	err := s.db.Table("views").First(&siteViews).Error
	if err != nil {
		fmt.Printf("cant get site views: %v \n", err)
		return
	}

	err = s.db.Unscoped().Table("views").Where("id = ?", siteViews.ID).Delete(&siteViews).Error
	if err != nil {
		fmt.Printf("cant delete site views: %v \n", err)
	}

	err = s.db.Table("views").Create(&models.SiteViews{Count: s.count}).Error
	if err != nil {
		fmt.Printf("cant write site views: %v \n", err)
	}
}

func (s *Server) getSiteViews() int {
	var siteViews models.SiteViews
	err := s.db.Table("views").First(&siteViews).Error
    fmt.Printf("got siteviews: %v %v \n", siteViews, err)
	if err != nil {
        fmt.Printf("creating initial views record")
        s.db.Table("views").Create(&models.SiteViews{Count: 0})
		return 0
	}
	return siteViews.Count
}
