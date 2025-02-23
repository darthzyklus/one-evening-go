package server

import (
	"sync"
)

type TweetsRepository interface {
	SaveTweet(tw Tweet) (int, error)
	Tweets() ([]Tweet, error)
}

type TweetsMemoryRepository struct {
	tweets []Tweet
	lock   sync.RWMutex
}

func (r *TweetsMemoryRepository) SaveTweet(tw Tweet) (int, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.tweets = append(r.tweets, tw)

	return len(r.tweets), nil
}

func (r *TweetsMemoryRepository) Tweets() ([]Tweet, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.tweets, nil
}
