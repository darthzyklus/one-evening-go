package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Server struct {
	TweetsRepository TweetsRepository
}

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type TweetResponse struct {
	ID int
}

func (s Server) Tweets(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	defer func() {
		duration := time.Since(start)

		fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
	}()

	if r.Method == http.MethodPost {
		s.AddTweet(w, r)
	} else if r.Method == http.MethodGet {
		s.ListTweets(w, r)
	}
}

func (s Server) AddTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	tw := Tweet{}

	if err := json.Unmarshal(body, &tw); err != nil {
		fmt.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if tw.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tweetID, err := s.TweetsRepository.SaveTweet(tw)

	if err != nil {
		fmt.Println("Failed to save tweet:", nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := TweetResponse{ID: tweetID}

	respJSON, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("Failed to marshal response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}

type tweetsList struct {
	Tweets []Tweet `json:"tweets"`
}

func (s Server) ListTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.TweetsRepository.Tweets()

	if err != nil {
		fmt.Println("Failed to get tweets:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := tweetsList{
		Tweets: tweets,
	}

	respJSON, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("Failed to marshal response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)

}
