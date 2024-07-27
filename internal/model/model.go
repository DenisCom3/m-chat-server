package model

type Chat struct {
	ID    int64    
	Name  string   
	UsersId []int64 
	IsActive bool
}

type Message struct {
	ChatId    int64 
	UserId    int64
	Message   string
}

type User struct {
	ID    int64
	Login string
}

type CreateChat struct {
	Users []User
	Name  string
}

type CreateMessage struct {
	ChatId    int64
	UserId    int64
	Text   string
}
