package service_test

import (
	"github.com/nrudolph/twitter/src/domain"
	"github.com/nrudolph/twitter/src/domain/user"
	"github.com/nrudolph/twitter/src/persistency"
	"github.com/nrudolph/twitter/src/service"
	"strings"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager(persistency.NewMemoryTweetWriter())
	userManager := service.NewUserManager()

	var tweet *domain.TextTweet
	u, _ := userManager.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe")
	text := "mi super tweet"
	tweet = domain.NewTextTweet(u, text)

	// Operation
	_, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweets()[0]

	if publishedTweet.GetUser().Username != u.Username && publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbout is %s: %s", u.Username, text, publishedTweet.GetUser().Username, publishedTweet.GetText())
	}
	if publishedTweet.GetDate() == nil {
		t.Error("Expeted date can't be nil")
	}
}

func TestWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager(persistency.NewMemoryTweetWriter())
	_ = service.NewUserManager()
	var tweet domain.Tweet

	var u *user.User
	text := "super tweet"

	tweet = domain.NewTextTweet(u, text)

	// Operation

	_, err := tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager(persistency.NewMemoryTweetWriter())
	userManager := service.NewUserManager()
	var tweet domain.Tweet

	u, _ := userManager.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe")
	var text string

	tweet = domain.NewTextTweet(u, text)

	// Operation

	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager(persistency.NewMemoryTweetWriter())
	userManager := service.NewUserManager()

	var tweet domain.Tweet

	u, _ := userManager.AddUser("pepe", "pepe@pepe.com", "pepe", "pepe")
	text := "super tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweetsuper tweettweetsuper tweet"

	tweet = domain.NewTextTweet(u, text)

	// Operation

	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "tweet over 140 characters" {
		t.Error("Expected error is user is required")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Init
	tweetManager := service.NewTweetManager(persistency.NewMemoryTweetWriter())
	userManager := service.NewUserManager()

	var tweet, secondTweet domain.Tweet

	tweetUser, _ := userManager.AddUser("pepe", "pepe", "pepe", "pepe")
	tweet = domain.NewTextTweet(tweetUser, "Mi primer tweet")
	secondTweet = domain.NewTextTweet(tweetUser, "Mi segundo tweet")

	// Operations
	_, _ = tweetManager.PublishTweet(tweet)
	_, _ = tweetManager.PublishTweet(secondTweet)

	// Validation

	publishedTweets := tweetManager.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishTweet := publishedTweets[0]
	secondPublishTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishTweet, tweetUser.Username, tweet.GetText()) {
		return
	}

	if !isValidTweet(t, secondPublishTweet, tweetUser.Username, secondTweet.GetText()) {
		return
	}
}

func isValidTweet(t *testing.T, tweet domain.Tweet, user string, text string) bool {
	if tweet.GetUser().Username != user {
		t.Errorf("Expected %s but was %s", user, tweet.GetUser().Username)
		return false
	}

	if tweet.GetText() != text {
		t.Errorf("Expected %s but was %s", text, tweet.GetText())
		return false
	}

	return true
}

func TestCanRetrieveTweetById(t *testing.T) {
	// Init
	tweetManager := service.NewTweetManager(persistency.NewMemoryTweetWriter())
	userManager := service.NewUserManager()
	var tweet domain.Tweet
	var id int

	tweetUser, _ := userManager.AddUser("pepe", "pepe@pepe.com", "Pepe", "cont")
	text := "Super tweet"

	tweet = domain.NewTextTweet(tweetUser, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet, _ := tweetManager.GetTweetById(id)
	isValidTweet(t, publishedTweet, tweetUser.Username, text)
}

func TestReturnsErrorWhenRetrieveTweetByIdInvalid(t *testing.T) {
	// Init
	tweetManager := service.NewTweetManager(persistency.NewMemoryTweetWriter())
	userManager := service.NewUserManager()
	var tweet domain.Tweet

	tweetUser, _ := userManager.AddUser("pepe", "pepe@pepe.com", "Pepe", "cont")
	text := "Super tweet"

	tweet = domain.NewTextTweet(tweetUser, text)

	// Operation
	_, _ = tweetManager.PublishTweet(tweet)

	// Validation
	_, err := tweetManager.GetTweetById(2)
	if err == nil {
		t.Error("Expected an error but was not found")
		return
	}
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// init
	tweetManager := service.NewTweetManager(persistency.NewMemoryTweetWriter())
	userManager := service.NewUserManager()
	var tweet, secondTweet, thirdTweet domain.Tweet

	user1, _ := userManager.AddUser("pepe", "pepe@pepe.com", "Pepe", "pepe")
	user2, _ := userManager.AddUser("Juan", "juan@pepe.com", "Juan", "pepe")

	text1 := "Primer tweet"
	text2 := "Segundo tweet"
	text3 := "Tercer tweet"

	tweet = domain.NewTextTweet(user1, text1)
	secondTweet = domain.NewTextTweet(user1, text2)
	thirdTweet = domain.NewTextTweet(user2, text3)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	//Operation
	count := tweetManager.CountTweetsByUser(user1)

	// Validation

	if count != 2 {
		t.Errorf("expected 2 but was found %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager(persistency.NewMemoryTweetWriter())
	userManager := service.NewUserManager()
	var tweet, secondTweet, thirdTweet domain.Tweet
	user1, _ := userManager.AddUser("pepe", "pepe@pepe.com", "Pepe", "pepe")
	anotherUser, _ := userManager.AddUser("nick", "nick@pepe.com", "Nick", "pepe")
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user1, text)
	secondTweet = domain.NewTextTweet(user1, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)
	// publish the 3 tweets

	userManager.LoginUser("pepe", "pepe")
	userManager.LoginUser("nick", "pepe")

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user1)

	// Validation
	if len(tweets) != 2 {
		t.Errorf("expected 2 but was found %d", len(tweets))
	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, user1.Username, text) {
		return
	}
	if !isValidTweet(t, secondPublishedTweet, user1.Username, secondText) {
		return
	}
}

func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {

	// Initialization
	var tweetWriter persistency.TweeterWriter
	tweetWriter = persistency.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet // Fill the tweet with data
	u := user.NewUser("p", "p", "p", "p")
	tweet = domain.NewTextTweet(u, "kjchsbkjac")
	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	memoryWriter := (tweetWriter).(*persistency.MemoryTweetWriter)
	savedTweet := memoryWriter.GetLastSavedTweet()

	if savedTweet == nil {
		t.Error("A tweet was expected but not found")
		return
	}
	if savedTweet.GetId() != id {
		t.Errorf("A tweet was expected wid id %d but was %d", id, savedTweet.GetId())
		return
	}
}

func TestCanSearchForTweetContainingText(t *testing.T) {
	// Initialization
	var tweetWriter persistency.TweeterWriter
	tweetWriter = persistency.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	// Create and publish a tweet

	newUser := user.NewUser("p", "p", "p", "p")
	tweet := domain.NewTextTweet(newUser, "My first tweet")
	tweetManager.PublishTweet(tweet)

	// Operation
	searchResult := make(chan domain.Tweet)
	query := "first"
	tweetManager.SearchTweetsContaining(query, searchResult)

	// Validation
	foundTweet := <-searchResult

	if foundTweet == nil {
		t.Error("Was expected to find one tweet but was nil")
	}
	if !strings.Contains(foundTweet.GetText(), query) {
		t.Errorf("Was expected that %s contains %s", foundTweet.GetText(), query)
	}
}
