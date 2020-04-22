package main

import (
	"task/social/facebook"
	"task/social/twitter"
	"task/social/linkedin"
	"task/social/exporter"
)
func main() {
	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkdin := new(linkedin.Linkedin)
	err := exporter.Export(twtr, "twtrdata.txt")
	err = exporter.Export(lnkdin, "lnkdindata.txt")
	err = exporter.Export(fb, "fbdata.txt")

	if err != nil {
		panic(err)
	}



	// twtr := new(twitter.Twitter)
	// lnkdin := new(linkedin.Linkedin)

	// ScrollFeeds(fb, twtr, lnkdin)
}

// ScrollFeeds prints all social media feeds



