package hap

type Service struct {
    Id int `json:"iid"`
    Type UUID `json:"type"`
    Characteristics []Characteristic `json:"characteristics"`
}