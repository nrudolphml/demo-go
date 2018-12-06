package domain

import (
	"github.com/nrudolph/twitter/src/domain/user"
	"time"
)

type Tweet interface {
	GetUser() *user.User
	GetText() string
	GetDate() *time.Time
	GetId() int
}
