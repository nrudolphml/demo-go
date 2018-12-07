package persistency

import (
	"github.com/nrudolph/twitter/src/domain"
	"os"
)

type FileTweetWriter struct {
	file *os.File
}

func (writer *FileTweetWriter) WriteTweet(tweet domain.Tweet) {
	go func() {
		if writer.file != nil {
			writer.file.WriteString(tweet.String() + "\n")
		}
	}()
}

func NewFileTweetWritter() *FileTweetWriter {
	file, _ := os.OpenFile(
		"pepe.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	return &FileTweetWriter{file}
}
