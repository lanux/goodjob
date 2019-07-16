package web

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/sessions"
	"time"
)

var Sessions *sessions.Sessions

func init() {
	// attach a session manager
	cookieName := "GOSESSIONID"
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey := []byte("the-big-and-secret-fash-key-here")
	blockKey := []byte("lot-secret-of-characters-big-too")
	secureCookie := securecookie.New(hashKey, blockKey)

	Sessions = sessions.New(sessions.Config{
		Cookie:  cookieName,
		Encode:  secureCookie.Encode,
		Decode:  secureCookie.Decode,
		Expires: time.Duration(30) * time.Minute,
	})
}
