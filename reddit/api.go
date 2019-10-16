package reddit

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	baseSubredditUrl = "https://reddit.com/r"
)

func getClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

func GetTopPosts(subreddit string) (error, *SubredditResponse) {
	logrus.Debug("Getting top Reddit posts.")
	client := getClient()
	// Call API.
	url := fmt.Sprintf("%s/%s/top/.json", baseSubredditUrl, subreddit)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err, nil
	}
	req.Header.Add("User-Agent", "RedditPlaylist github.com/farzaa/RedditPlaylist")
	resp, err := client.Do(req)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	subredditResponse := &SubredditResponse{}
	if err := json.NewDecoder(resp.Body).Decode(subredditResponse); err != nil {
		return err, nil
	}

	return nil, subredditResponse
}
