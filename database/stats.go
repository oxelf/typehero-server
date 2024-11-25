package database

import "typehero_server/models"

func (db *Database) GetStats() (*models.Stats, error) {
	var count int64
	err := db.Table("results").Count(&count).Error
	if err != nil {
		return nil, err
	}
	return &models.Stats{TestsStarted: int(count)}, nil
}
