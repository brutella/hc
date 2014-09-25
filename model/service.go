package model

type Service struct {
    Id int `json:"iid"`
    Type ServiceType `json:"type"`
    Characteristics []Characteristic `json:"characteristics"`
}