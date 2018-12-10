package main

import (
	"github.com/abiosoft/ishell"
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/persistency"
	"github.com/nrudolph/twitter/src/rest"
	"github.com/nrudolph/twitter/src/service"
	"strconv"
)

func main() {
	tweetManager := service.NewTweetManager(persistency.NewFileTweetWritter())
	userManager := service.NewUserManager()

	restServer := rest.NewRestServer(tweetManager, userManager)
	restServer.Start()

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the type of tweet to publish (text/image/quote): ")

			tweetType := c.ReadLine()

			switch tweetType {
			case "text":
				publishTextTweet(c, userManager, tweetManager)
			case "image":
				publishImageTweet(c, userManager, tweetManager)
			case "quote":
				publishQuoteTweet(c, userManager, tweetManager)
			default:
				c.Println("The type of tweet is not valid\n")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "deleteTweet",
		Help: "Deletes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username/email/nickname: ")

			identifier := c.ReadLine()

			c.Print("Write your tweets id: ")

			id := c.ReadLine()

			user, err := userManager.GetUser(identifier)

			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			if !userManager.IsUserLoggedIn(user) {
				c.Println("The user must login to publish tweets\n")
				return
			}

			numId, err := strconv.Atoi(id)

			if err != nil {
				c.Println("The id is not valid\n")
				return
			}
			_, err = tweetManager.DeleteTweet(user, numId)

			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			c.Print("Tweet deleted\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsWithId",
		Help: "Shows all tweets with Id and date",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetManager.GetTweets()

			for _, tweet := range tweets {
				c.Printf("%s\n", tweet.FullString())
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetManager.GetTweets()

			for _, tweet := range tweets {
				c.Printf("%s\n", tweet.String())
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "searchTweet",
		Help: "Searchs a tweet that contains the query",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your query: ")
			query := c.ReadLine()

			searchResult := make(chan domain.Tweet)
			tweetManager.SearchTweetsContaining(query, searchResult)

			amountOfResults := 0
			for tweet := range searchResult {
				if tweet == nil {
					break
				}
				amountOfResults++
				c.Printf("%s\n", tweet.String())
			}

			c.Printf("Found %d result/s\n", amountOfResults)
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

			if _, err := userManager.AddUser(username, email, nickname, pass); err != nil {
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

			tweet, err := tweetManager.GetTweetById(numId)
			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			c.Printf("%s\n", tweet.FullString())

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

			owner, err := userManager.GetUser(identifier)

			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			count := tweetManager.CountTweetsByUser(owner)

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

			if _, err := userManager.LoginUser(identifier, pass); err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			c.Print("User loggedIn\n")

			return

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "logout",
		Help: "Log out a user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write the user's username/email/nickname: ")
			identifier := c.ReadLine()

			owner, err := userManager.GetUser(identifier)

			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			if !userManager.LogoutUser(owner) {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			c.Print("User logged out\n")

			return

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "followUser",
		Help: "Adds user to my followers list",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write the user's username/email/nickname: ")
			identifier := c.ReadLine()

			owner, err := userManager.GetUser(identifier)

			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			c.Print("Write the user to follow username/email/nickname: ")
			identifierToFollow := c.ReadLine()

			userToFollow, err := userManager.GetUser(identifierToFollow)

			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			err = userManager.FollowUser(owner, userToFollow)

			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			c.Printf("User %s is now following %s\n", owner.Nickname, userToFollow.Nickname)

			return

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showUserFollowers",
		Help: "List the users that the user follows",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write the user's username/email/nickname: ")
			identifier := c.ReadLine()

			owner, err := userManager.GetUser(identifier)

			if err != nil {
				c.Println("An error has occurred: ", err, "\n")
				return
			}

			followers := userManager.GetUserFollowers(owner)

			if followers == nil {
				c.Printf("%s doesn't follow other users \n", owner.Nickname)
				return
			}

			for _, v := range followers {
				c.Printf("User: %s \n", v.Nickname)
			}

			c.Printf("Following %d users\n", len(followers))

			return

		},
	})

	shell.Run()
}

func publishTextTweet(c *ishell.Context, userManager *service.UserManager, tweetManager *service.TweetManager) {
	c.Print("Write your username: ")

	username := c.ReadLine()

	c.Print("Write your tweet: ")

	text := c.ReadLine()

	user, err := userManager.GetUser(username)

	if err != nil {
		c.Println("An error has occurred: ", err, "\n")
		return
	}

	if !userManager.IsUserLoggedIn(user) {
		c.Println("The user must login to publish tweets\n")
		return
	}

	tweet := domain.NewTextTweet(user, text)

	if _, err := tweetManager.PublishTweet(tweet); err != nil {
		c.Println("An error has occurred: ", err, "\n")
		return
	}

	c.Print("Tweet sent\n")
	return
}

func publishImageTweet(c *ishell.Context, userManager *service.UserManager, tweetManager *service.TweetManager) {
	c.Print("Write your username: ")

	username := c.ReadLine()

	c.Print("Write your tweet: ")

	text := c.ReadLine()

	c.Print("Write the image URL: ")

	url := c.ReadLine()

	user, err := userManager.GetUser(username)

	if err != nil {
		c.Println("An error has occurred: ", err, "\n")
		return
	}

	if !userManager.IsUserLoggedIn(user) {
		c.Println("The user must login to publish tweets\n")
		return
	}

	tweet := domain.NewImageTweet(user, text, url)

	if _, err := tweetManager.PublishTweet(tweet); err != nil {
		c.Println("An error has occurred: ", err, "\n")
		return
	}

	c.Print("Tweet sent\n")
	return
}

func publishQuoteTweet(c *ishell.Context, userManager *service.UserManager, tweetManager *service.TweetManager) {
	c.Print("Write your username: ")

	username := c.ReadLine()

	c.Print("Write your tweet: ")

	text := c.ReadLine()

	c.Print("Write the quoted tweet id: ")

	id := c.ReadLine()

	numId, err := strconv.Atoi(id)

	if err != nil {
		c.Println("The id is not valid\n")
		return
	}

	quotedTweet, err := tweetManager.GetTweetById(numId)

	if err != nil {
		c.Println("An error has occurred: ", err, "\n")
		return
	}

	user, err := userManager.GetUser(username)

	if err != nil {
		c.Println("An error has occurred: ", err, "\n")
		return
	}

	if !userManager.IsUserLoggedIn(user) {
		c.Println("The user must login to publish tweets\n")
		return
	}

	tweet := domain.NewQuoteTweet(user, text, quotedTweet)

	if _, err := tweetManager.PublishTweet(tweet); err != nil {
		c.Println("An error has occurred: ", err, "\n")
		return
	}

	c.Print("Tweet sent\n")

	return
}
