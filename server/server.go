package server

import (
	"fmt"
	"sync"
	"time"
	"typehero_server/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

type Server struct {
	db         *database.Database
	countMutex sync.Mutex
	count      int
}

func StartServer(db *database.Database) error {
	s := Server{
		db:         db,
		count:      0,
		countMutex: sync.Mutex{},
	}

	count := s.getSiteViews()
	s.count = count

	viewsTicker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			<-viewsTicker.C
			s.saveSiteViews()
		}
	}()

	router := gin.New()
	log := logrus.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowAllOrigins = true

	router.Use(ginlogrus.Logger(log), cors.New(config), gin.Recovery())

	router.GET("/healthcheck", func(c *gin.Context) {
		fmt.Println("got healthcheck")
		c.Status(200)
	})

	router.POST("result", s.postResults)
	router.GET("leaderboard", s.getLeaderboard)
	router.GET("stats", s.getStats)
	router.GET("view", s.increaseSiteViews)

	return router.Run(":8011")
}
