// Parking Backend - Runtime
//
// Session token support.
// ATM, only support sessions stored in RAM.
//
// 2015

package runtime

import (
	"math/rand"
	"net/http"
	"time"

	"bitbucket.org/remeh/parking/db/model"
	. "bitbucket.org/remeh/parking/logger"
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
	Get(string) (Session, bool)
	GetFromRequest(*http.Request) (Session, bool)
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

	Debug(len(s.sessions), "session tokens in RAM.")

	return session
}

func (s RAMSessionStorage) Get(token string) (Session, bool) {
	session := s.sessions[token]
	return session, len(session.User.Uid) > 0
}

func (s RAMSessionStorage) GetFromRequest(request *http.Request) (Session, bool) {
	c, err := request.Cookie(COOKIE_TOKEN_KEY)

	if err == nil && c != nil {
		session := s.sessions[c.Value]
		return session, len(session.User.Uid) > 0
	}

	return Session{}, false
}

// TODO(remy): expiration methods.

func randomToken(size int) string {
	result := make([]rune, size)
	for i := range result {
		result[i] = runes[rand.Intn(len(runes))]
	}
	return string(result)
}
