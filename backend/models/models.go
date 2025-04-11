package models


import (
	//"fmt"
)



type Project struct {
	ProjectID   int     `gorm:"primaryKey;autoIncrement" json:"project_id"`
	Name        string  `gorm:"not null" json:"name"`
	DateRange   string  `gorm:"not null" json:"name"`
	ImageURL    string  `gorm:"not null json:"image_url"` 
	GithubURL   string  `gorm:"not null json:"github_url"` 
	Description string  `gorm:"not null json:"description"` 
}



type Image struct {

}