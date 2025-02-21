package server

type TweetsRepository interface {
	SaveTweet(tw Tweet) (int, error)
	Tweets() ([]Tweet, error)
}

type TweetsMemoryRepository struct {
	tweets []Tweet
}

func (r *TweetsMemoryRepository) SaveTweet(tw Tweet) (int, error) {
	r.tweets = append(r.tweets, tw)

	return len(r.tweets), nil
}

func (r *TweetsMemoryRepository) Tweets() ([]Tweet, error) {
	return r.tweets, nil
}
