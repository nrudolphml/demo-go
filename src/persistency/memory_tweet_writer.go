package persistency

import "github.com/nrudolph/twitter/src/domain"

type MemoryTweetWriter struct {
	tweets []domain.Tweet
}

func (writer *MemoryTweetWriter) WriteTweet(tweet domain.Tweet) {
	writer.tweets = append(writer.tweets, tweet)
}

func (writer *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	return writer.tweets[len(writer.tweets)-1]
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	return &MemoryTweetWriter{make([]domain.Tweet, 0)}
}
