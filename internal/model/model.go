package model

type Chat struct {
	ID    int64    
	Name  string   
	Users []string 
}

type Message struct {
	ChatName  string
	UserLogin string
	Message   string
}

