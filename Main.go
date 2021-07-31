package main

import (
	"fmt"
	"github.com/jzelinskie/geddit"
	"log"
)

func main() {
	o, err := geddit.NewOAuthSession(
		"njAZY90LDrZu-g",
		"2cI1QZEoocLeipA6OpRAKKlpJTJIJQ",
		"Testing Geddit Bot with OAuth v0.1 by u/aggrolite - see source @ github.com/aggrolite/geddit/master",
		"",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Login using our personal reddit account.
	err = o.LoginAuth("TitlePullerHomework", "pullingtitles")
	if err != nil {
		log.Fatal(err)
	}

	// We can pass options to the query if desired (blank for now).
	opts := geddit.ListingOptions{}

	// Fetch posts from r/videos, sorted by Hot.
	posts, err := o.SubredditSubmissions("WallStreetBets", geddit.HotSubmissions, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Print the title and URL of each post that has a youtube.com domain.
	for _, p := range posts {
		if p.LinkFlairText == "Gain" {
		fmt.Printf("%s (%d) - %s\n", p.Title, p.Score, p.URL)
		}
	}
}