package models

import "gorm.io/gorm"

type SiteViews struct {
    gorm.Model
	Count int `json:"count"`
}
