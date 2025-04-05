package models

type Car struct {
	ID    uint   `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  uint   `json:"year"`
	Price uint   `json:"price"`
}
