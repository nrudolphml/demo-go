package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/service"
	"net/http"
	"strconv"
)

type RestServer struct {
	tweetManager *service.TweetManager
	userManager  *service.UserManager
}

func NewRestServer(tweetManager *service.TweetManager, userManager *service.UserManager) *RestServer {
	restServer := RestServer{tweetManager, userManager}
	restServer.Start()
	return &restServer
}

func (restServer *RestServer) Start() {
	router := gin.Default()

	tweets := router.Group("/tweets")
	{
		tweets.GET("", restServer.listTweets)
		tweets.GET("/user/:identification", restServer.listUserTweets)
		tweets.GET("/search", restServer.searchTweets)
	}
	tweet := router.Group("/tweet")
	{
		tweet.GET("/:id", restServer.viewTweet)
		tweet.POST("/text", restServer.Authenticator(), restServer.publishTextTweet)
		tweet.POST("/image", restServer.Authenticator(), restServer.publishImageTweet)
		tweet.POST("/quote", restServer.Authenticator(), restServer.publishQuoteTweet)
	}
	router.POST("/user", restServer.registerUser)
	router.POST("/login", restServer.login)
	router.POST("/logout", restServer.Authenticator(), restServer.logout)

	go router.Run()
}

func (restServer *RestServer) listTweets(c *gin.Context) {
	c.JSON(http.StatusOK, restServer.tweetManager.GetTweets())
}

func (restServer *RestServer) viewTweet(c *gin.Context) {
	paramId := c.Param("id")
	id, e := strconv.Atoi(paramId)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	tweet, e := restServer.tweetManager.GetTweetById(id)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	c.JSON(http.StatusOK, tweet)

}

func (restServer *RestServer) listUserTweets(c *gin.Context) {
	identification := c.Param("identification")

	owner, e := restServer.userManager.GetUser(identification)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	c.JSON(http.StatusOK, restServer.tweetManager.GetTweetsByUser(owner))
}

func (restServer *RestServer) registerUser(c *gin.Context) {
	var userToRegister RegisterUser
	if err := c.ShouldBind(&userToRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, e := restServer.userManager.AddUser(userToRegister.Username, userToRegister.Email, userToRegister.Nickname, userToRegister.Password)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user register"})
}

func (restServer *RestServer) login(c *gin.Context) {
	var loginInfo Login
	if err := c.ShouldBind(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, e := restServer.userManager.LoginUser(loginInfo.Identification, loginInfo.Password)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	if result {
		c.JSON(http.StatusOK, gin.H{"status": "user logged in"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "user could not be logged in"})
	}
}

func (restServer *RestServer) logout(c *gin.Context) {
	var logoutInfo Login
	if err := c.ShouldBind(&logoutInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userToLogout, e := restServer.userManager.GetUser(logoutInfo.Identification)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	result := restServer.userManager.LogoutUser(userToLogout)

	if result {
		c.JSON(http.StatusOK, gin.H{"status": "user logged out"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "user could not be logged out"})
	}
}

func (restServer *RestServer) publishTextTweet(c *gin.Context) {
	var tweet TweetRaw
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	owner, e := restServer.userManager.GetUser(tweet.UserIdentification)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	textTweet := domain.NewTextTweet(owner, tweet.Text)
	i, e := restServer.tweetManager.PublishTweet(textTweet)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "tweet was sent with id: " + strconv.Itoa(i)})
}

func (restServer *RestServer) publishImageTweet(c *gin.Context) {
	var tweet TweetRaw
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	owner, e := restServer.userManager.GetUser(tweet.UserIdentification)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	textTweet := domain.NewImageTweet(owner, tweet.Text, tweet.Url)
	i, e := restServer.tweetManager.PublishTweet(textTweet)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "tweet was sent with id: " + strconv.Itoa(i)})
}

func (restServer *RestServer) publishQuoteTweet(c *gin.Context) {
	var tweet TweetRaw
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	owner, e := restServer.userManager.GetUser(tweet.UserIdentification)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	quotedTweet, e := restServer.tweetManager.GetTweetById(tweet.QuoteId)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	textTweet := domain.NewQuoteTweet(owner, tweet.Text, quotedTweet)
	i, e := restServer.tweetManager.PublishTweet(textTweet)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "tweet was sent with id: " + strconv.Itoa(i)})
}

func (restServer *RestServer) searchTweets(c *gin.Context) {
	query := c.Query("q")

	searchResult := make(chan domain.Tweet)
	restServer.tweetManager.SearchTweetsContaining(query, searchResult)

	results := make([]domain.Tweet, 0)
	for tweet := range searchResult {
		if tweet == nil {
			break
		}
		results = append(results, tweet)
	}

	c.JSON(http.StatusOK, results)
}

func (restServer *RestServer) Authenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdentification := c.GetHeader("UserIdentification")
		if userIdentification == "" {
			c.JSON(http.StatusUnauthorized, "The user is not log in")
			c.Abort()
			return
		}
		user, e := restServer.userManager.GetUser(userIdentification)

		if e != nil {
			c.JSON(http.StatusUnauthorized, "The user does not exist")
			c.Abort()
			return
		}

		if !restServer.userManager.IsUserLoggedIn(user) {
			c.JSON(http.StatusUnauthorized, "The user is not log in")
			c.Abort()
			return
		}
		return
	}
}
