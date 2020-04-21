package exporter

import (
	"fmt"
	"os"
	"task"
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
