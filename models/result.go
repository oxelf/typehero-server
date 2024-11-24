package models

import "time"

type Result struct {
    UserName string `json:"userName"`
    UserId string `json:"userId"`
    Date time.Time `json:"date"`
    WPM float64 `json:"wpm"`
    Accurary float64 `json:"accuracy"`
    TimeSeconds float64 `json:"timeSeconds"`
    Mode string `json:"mode"`
    Language string `json:"language"`
    WordAmount int `json:"wordAmount"`
    Rank int `json:"rank" gorm:"rank"`
}
