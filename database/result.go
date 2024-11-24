package database

import "typehero_server/models"

func (db *Database) CreateResult(result models.Result) error  {
    err := db.Table("results").Create(&result).Error
    return err
}
