package model

type Compareable interface {
    Equal(interface{}) bool
}