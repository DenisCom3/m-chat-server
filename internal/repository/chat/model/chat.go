package model

type Chat struct {
	ID    int64
	Name  string
	Users []int64
	IsActive bool
}