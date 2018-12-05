package service

var tweet string

func PublishTweet(tweetToPublish string) {
	tweet = tweetToPublish
}

func GetTweet() string {
	return tweet
}
