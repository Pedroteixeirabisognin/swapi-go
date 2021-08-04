package model

type Planet struct {
	ID      string `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Ground  string `json:"ground"`
}
