package persistency

import "github.com/nrudolph/twitter/src/domain"

type TweeterWriter interface {
	WriteTweet(tweet domain.Tweet)
}
