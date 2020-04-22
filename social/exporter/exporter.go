package exporter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"task"
	"gopkg.in/yaml.v3"
)

//ScrollFeeds seeds the feeds of social media and prints them out
func ScrollFeeds(platforms ...task.SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("=================================")
	}
}

//Export is used in main
func Export(u task.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	for _, fd := range u.Feed() {
		n, err := f.WriteString(fd + "\n")
		if err != nil {
			return err
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}
//JSON exports .json
func JSON(u task.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer f.Close()

	if err != nil {
		return err
	}

	for _, jd := range u.Feed() {
		o, err := json.MarshalIndent(jd, "\n", "\n")
		if err != nil {
			return err
		}

		bytesWritten, err := f.Write(o)

		if err != nil {
			return err
		}
		fmt.Printf("wrote %d bytes\n", bytesWritten)
	}

	return nil
}

// XML prints the social media feed in .xml
func XML(u task.SocialMedia, filename string) error {
	x, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer x.Close()

	if err != nil {
		return err
	}

	for _, xd := range u.Feed() {
		f, err := xml.MarshalIndent(xd, "\n", "\n")
		if err != nil {
			return err
		}

		bytesWritten, err := x.Write(f)

		if err != nil {
			return err
		}
		fmt.Printf("wrote %d bytes\n", bytesWritten)
	}
	return nil
}

// Yaml exports in a yaml file
func Yaml(u task.SocialMedia, filename string) error {
	q, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer q.Close()

	if err != nil {
		return err
	}

	for _, yd := range u.Feed() {
		f, err := yaml.Marshal(yd)
		if err != nil {
			return err
		}

		bytesWritten, err := q.Write(f)

		if err != nil {
			return err
		}
		fmt.Printf("wrote %d bytes\n", bytesWritten)
	}
	return nil
}