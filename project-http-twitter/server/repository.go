package server

type TweetsRepository interface {
	AddTweet(t Tweet) (int, error)
	Tweets() ([]Tweet, error)
}

type TweetsMemoryRepository struct {
	tweets []Tweet
}

func (repository *TweetsMemoryRepository) AddTweet(t Tweet) (int, error) {
	repository.tweets = append(repository.tweets, t)

	return len(repository.tweets), nil
}

func (repository *TweetsMemoryRepository) Tweets() ([]Tweet, error) {
	tweets := repository.tweets

	return tweets, nil
}
