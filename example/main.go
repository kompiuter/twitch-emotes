package main

import (
	"fmt"
	"log"

	twem "github.com/kompiuter/twitch-emotes"
)

func main() {
	// TODO: Make example prettier
	gl, err := twem.Global()
	if err != nil {
		log.Fatal(err)
	}
	for i, e := range gl.Emotes {
		if i > 5 {
			break
		}
		fmt.Println(e.Code, e.Description, e.FirstSeen, e.ImageID)
	}

	sub, err := twem.Subscriber()
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range sub.Channels {
		if c.Title == "LIRIK" {
			fmt.Println(c.Badge)
			fmt.Println(c.Badge3m)
			fmt.Println(c.Badge6m)
			fmt.Println(c.Badge12m)
			fmt.Println(c.Badge24m)
			fmt.Println(c.BadgeStarting)
			fmt.Println(c.ChannelID)
			fmt.Println(c.Desc)
			fmt.Println(c.FirstSeen)
			fmt.Println(c.ID)
			fmt.Println(c.Link)
			fmt.Println(c.Set)
			fmt.Println(c.Title)
			for _, e := range c.Emotes {
				fmt.Println(e.Code, e.ImageID)
			}
		}
	}
}
