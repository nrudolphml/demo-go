package persistency

import "github.com/nrudolph/twitter/src/domain"

type TweeterWritter interface {
	WriteTweet(tweet domain.Tweet)
}
