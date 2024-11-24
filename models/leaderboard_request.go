package models

type LeaderboardRequest struct {
    Mode string `json:"mode"`
    WordAmount int `json:"wordAmount"`
    Language string `json:"language"`
    Page int `json:"page"`
}
