package database

import (
	"time"
	"typehero_server/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
    *gorm.DB
}

func InitDatabase() (*Database, error)  {
db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{
		Logger: logger.New(
			log.StandardLogger(),
			logger.Config{
				SlowThreshold:             time.Second * 5, // Slow SQL threshold
				LogLevel:                  logger.Silent,   // Log level
				IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,           // Don't include params in the SQL log
				Colorful:                  false,           // Disable color
			},
		),
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.Table("results").AutoMigrate(&models.Result{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
return &Database{db}, nil
}
