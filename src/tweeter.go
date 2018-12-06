package main

import (
	"github.com/abiosoft/ishell"
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")

			username := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			user, err := service.GetUser(username)

			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			tweet := domain.NewTweet(user, text)

			if err := service.PublishTweet(tweet); err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()

			for _, tweet := range tweets {
				c.Printf("%s: %s, %d-%02d-%02d %02d:%02d\n", tweet.User.Nickname, tweet.Text,
					(*tweet.Date).Year(), (*tweet.Date).Month(), (*tweet.Date).Day(), (*tweet.Date).Hour(),
					(*tweet.Date).Minute())
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "registerUser",
		Help: "Registers a user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write your username: ")
			username := c.ReadLine()
			c.Print("Write your email: ")
			email := c.ReadLine()
			c.Print("Write your nickname: ")
			nickname := c.ReadLine()
			c.Print("Write your password: ")
			pass := c.ReadLine()

			if _, err := service.AddUser(username, email, nickname, pass); err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			c.Print("User added\n")

			return

		},
	})

	shell.Run()

}
