package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Setup Twitter API credentials from environment variables
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	// Check all credentials are available
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatalf("Missing Twitter API credentials. Please set the environment variables.")
	}

	// Post a new tweet
	tweetText := "Hello from Twitter API using only net/http!"
	err := postTweet(consumerKey, consumerSecret, accessToken, accessSecret, tweetText)
	if err != nil {
		log.Fatalf("Error posting tweet: %v", err)
	}

	// Delete a tweet (using a valid tweet ID)
	tweetID := "your_tweet_id_here" // Replace with actual tweet ID
	err = deleteTweet(consumerKey, consumerSecret, accessToken, accessSecret, tweetID)
	if err != nil {
		log.Fatalf("Error deleting tweet: %v", err)
	}
}

// postTweet function
func postTweet(consumerKey, consumerSecret, accessToken, accessSecret, tweetText string) error {
	tweetURL := "https://api.twitter.com/1.1/statuses/update.json"
	data := url.Values{}
	data.Set("status", tweetText)

	req, err := http.NewRequest("POST", tweetURL, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", "OAuth ...") // Placeholder for the actual OAuth header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to post tweet: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to post tweet, status code: %d", resp.StatusCode)
	}

	fmt.Println("Tweet posted successfully!")
	return nil
}

// deleteTweet function
func deleteTweet(consumerKey, consumerSecret, accessToken, accessSecret, tweetID string) error {
	deleteURL := fmt.Sprintf("https://api.twitter.com/1.1/statuses/destroy/%s.json", tweetID)

	req, err := http.NewRequest("POST", deleteURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", "OAuth ...") // Placeholder for the actual OAuth header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete tweet: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete tweet, status code: %d", resp.StatusCode)
	}

	fmt.Println("Tweet deleted successfully!")
	return nil
}
