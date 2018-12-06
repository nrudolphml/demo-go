package main

import (
	"github.com/abiosoft/ishell"
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/service"
	"strconv"
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

			if !service.IsUserLoggedIn(user) {
				c.Println("The user must login to publish tweets\n")
				return
			}

			tweet := domain.NewTweet(user, text)

			if _, err := service.PublishTweet(tweet); err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()

			for _, tweet := range tweets {
				printTweet(c, tweet)
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

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetWithId",
		Help: "Get tweet by Id",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write tweet id: ")
			id := c.ReadLine()

			numId, nanId := strconv.Atoi(id)

			if nanId != nil {
				c.Println("An error has occurred: ", nanId, "\n")
				return
			}

			tweet, err := service.GetTweetById(numId)
			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			printTweet(c, tweet)

			c.Print("User added\n")

			return

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countUserTweets",
		Help: "Counts tweets publish by user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write user's username/email/nickname: ")
			identifier := c.ReadLine()

			owner, err := service.GetUser(identifier)

			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			count := service.CountTweetsByUser(owner)

			c.Printf("Amount of tweets by %s: %d\nhel", owner.Nickname, count)

			c.Print("User loggedIn\n")

			return

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "Log in a user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write your username/email/nickname: ")
			identifier := c.ReadLine()
			c.Print("Write your password: ")
			pass := c.ReadLine()

			if _, err := service.LoginUser(identifier, pass); err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			c.Print("User loggedIn\n")

			return

		},
	})

	shell.Run()

}

func printTweet(c *ishell.Context, tweet *domain.Tweet) {
	c.Printf("%s: %s, %d-%02d-%02d %02d:%02d\n", tweet.User.Nickname, tweet.Text,
		(*tweet.Date).Year(), (*tweet.Date).Month(), (*tweet.Date).Day(), (*tweet.Date).Hour(),
		(*tweet.Date).Minute())
}
