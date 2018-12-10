package rest

import "github.com/nrudolph/twitter/src/domain/user"

type RegisterUser struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Nickname string `json:"Nickname"`
}

func (ru *RegisterUser) ToUser() *user.User {
	return user.NewUser(ru.Username, ru.Email, ru.Password, ru.Nickname)
}

type Login struct {
	Identification string `json:"Identification"`
	Password       string `json:"Password"`
}

type TweetRaw struct {
	UserIdentification string `json:"UserIdentification"`
	Text               string `json:"Text"`
	Url                string `json: "url"`
	QuoteId            int    `json:"quoteId"`
}
