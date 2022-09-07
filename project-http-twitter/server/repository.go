package server

import "sync"

type TweetsRepository interface {
	AddTweet(t Tweet) (int, error)
	Tweets() ([]Tweet, error)
}

type TweetsMemoryRepository struct {
	tweets []Tweet
	lock   sync.RWMutex
}

func (repository *TweetsMemoryRepository) AddTweet(t Tweet) (int, error) {
	repository.lock.Lock()
	defer repository.lock.Unlock()

	repository.tweets = append(repository.tweets, t)

	return len(repository.tweets), nil
}

func (repository *TweetsMemoryRepository) Tweets() ([]Tweet, error) {
	repository.lock.RLock()
	defer repository.lock.RUnlock()

	tweets := repository.tweets

	return tweets, nil
}
