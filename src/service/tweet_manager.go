package service

import "github.com/nrudolph/twitter/src/domain"

var tweet *domain.Tweet

func PublishTweet(tweetToPublish *domain.Tweet) {
	tweet = tweetToPublish
}

func GetTweet() *domain.Tweet {
	return tweet
}
