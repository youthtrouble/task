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

    j := exporter.Export(twtr, "twtrdata.json")
	j = exporter.Export(lnkdin, "lnkdindata.json")
	j = exporter.Export(fb, "fbdata.json")

	if j != nil {
		panic(err)
	}

	x := exporter.Export(twtr, "twtrdata.xml")
	x = exporter.Export(lnkdin, "lnkdindata.xml")
	x = exporter.Export(fb, "fbdata.xml")

	if x != nil {
		panic(err)
	}
	

	y := exporter.Export(twtr, "twtrdata.yaml")
	y = exporter.Export(lnkdin, "lnkdindata.yaml")
	y = exporter.Export(fb, "fbdata.yaml")

	if y != nil {
		panic(err)
	}
}





