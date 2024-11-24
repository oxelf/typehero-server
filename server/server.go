package server

import (
	"fmt"
	"typehero_server/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

type Server struct {
db *database.Database
}


func StartServer(db *database.Database) error {
    s := Server{
        db: db,
    }

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
    } )

    router.POST("result", s.postResults)
    router.GET("leaderboard", s.getLeaderboard)

    return router.Run(":8011")
}

