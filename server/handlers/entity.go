package handlers

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type PublicUser struct{
	Id int `json:"id"`
	Login    string `json:"login"`
}