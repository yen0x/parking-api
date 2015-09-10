// Parking Backend - Runtime
//
// Session token support.
// ATM, only support sessions stored in RAM.
//
// 2015

package runtime

import (
	"math/rand"
	"time"

	"bitbucket.org/remeh/parking/db/model"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")

const (
	TOKEN_SIZE = 32
)

type Session struct {
	Token        string
	CreationTime time.Time
	User         model.User
}

type SessionStorage interface {
	New(model.User) Session
	Get(string) Session
}

// Sessions stored in RAM.
type RAMSessionStorage struct {
	sessions map[string]Session
}

func CreateRAMSessionStorage() RAMSessionStorage {
	return RAMSessionStorage{
		sessions: make(map[string]Session),
	}
}

func (s *RAMSessionStorage) New(user model.User) Session {
	session := Session{
		Token:        randomToken(TOKEN_SIZE),
		CreationTime: time.Now(),
		User:         user,
	}

	s.sessions[session.Token] = session
	return session
}

func (s RAMSessionStorage) Get(token string) Session {
	return s.sessions[token]
}

// TODO(remy): expiration methods.

func randomToken(size int) string {
	result := make([]rune, size)
	for i := range result {
		result[i] = runes[rand.Intn(len(runes))]
	}
	return string(result)
}
