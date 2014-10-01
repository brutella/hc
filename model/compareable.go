package model

type Compareable interface {
    Equal(other interface{}) bool
}