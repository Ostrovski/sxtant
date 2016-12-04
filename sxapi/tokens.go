package sxapi

import (
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

type AccessToken struct {
	CreatedAt time.Time
	ExpiredAt time.Time
	UserId    int
	Token     string
	Quota     int
}

type TokenPool interface {
	GetAny() (*AccessToken, error)
	GetByUser(userId int) (*AccessToken, error)
	PutBack(*AccessToken) error
}

type TokenRegistry interface {
	Register(token *AccessToken) error
	Unregister(token *AccessToken) error
}

type TokenHolder struct {
	sync.Mutex
	mongoSes *mgo.Session
}

func NewTokenHolder(mongoSes *mgo.Session) *TokenHolder {

}
