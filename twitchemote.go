/*Package twitchemote is a wrapper for the unofficial Twitch API's found at
---> https://twitchemotes.com/apidocs <---
Functions in this package return information pertaining to Twitch global emotes
and Twitch emotes obtained from subscriber channels*/
package twitchemote

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Subscriber calls the subscriber emoticon API and returns a *SubscriberResult
// which contains all subscriber channels with their emoticons
func Subscriber() (*SubscriberResult, error) {
	// subs is the API for retrieving all Subscriber emoticons
	const subsAPI = "https://twitchemotes.com/api_cache/v2/subscriber.json"

	resp, err := http.Get(subsAPI)
	if err != nil {
		return nil, fmt.Errorf("subscriber: %v", err)
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	data := struct {
		Meta struct {
			GeneratedAt string `json:"generated_at"`
		} `json:"meta"`
		Template ImageTemplate `json:"template"`
		Channels map[string]struct {
			Badge         string     `json:"badge"`
			Badge3m       string     `json:"badge_3m"`
			Badge6m       string     `json:"badge_6m"`
			Badge12m      string     `json:"badge_12m"`
			Badge24m      string     `json:"badge_24m"`
			BadgeStarting string     `json:"badge_starting"`
			ChannelID     string     `json:"channel_id"`
			Desc          string     `json:"desc"`
			Emotes        []Emoticon `json:"emotes"`
			FirstSeen     string     `json:"first_seen"`
			ID            string     `json:"id"`
			Link          string     `json:"link"`
			Set           int        `json:"set"`
			Title         string     `json:"title"`
		} `json:"channels"`
	}{}
	if err := dec.Decode(&data); err != nil {
		return nil, fmt.Errorf("subscriber: %v", err)
	}
	t, err := time.Parse(time.RFC3339, data.Meta.GeneratedAt)
	if err != nil {
		log.Printf("could not parse 'generated at' time %q: %v\n", data.Meta.GeneratedAt, err)
	}
	res := SubscriberResult{
		Meta:     Metadata{GeneratedAt: t},
		Template: data.Template,
		Channels: make([]Channel, 0, len(data.Channels)),
	}
	for _, v := range data.Channels {
		t := time.Time{}
		if v.FirstSeen != "" {
			t, err = time.Parse("2006-01-02 15:04:05", v.FirstSeen)
			if err != nil {
				log.Printf("could not parse time %q of channel %q: %v\n", v.FirstSeen, v.Title, err)
			}
		}
		res.Channels = append(res.Channels, Channel{
			Badge:         v.Badge,
			Badge3m:       v.Badge3m,
			Badge6m:       v.Badge6m,
			Badge12m:      v.Badge12m,
			Badge24m:      v.Badge24m,
			BadgeStarting: v.BadgeStarting,
			ChannelID:     v.ChannelID,
			Desc:          v.Desc,
			Emotes:        v.Emotes,
			FirstSeen:     t,
			ID:            v.ID,
			Link:          v.Link,
			Set:           v.Set,
			Title:         v.Title,
		})
	}
	return &res, nil
}

// Global calls the global emoticon API and returns a *GlobalResult which contains all global emoticons
func Global() (*GlobalResult, error) {
	// global is the API for retrieving all Global emoticons
	const globalAPI = "https://twitchemotes.com/api_cache/v2/global.json"

	resp, err := http.Get(globalAPI)
	if err != nil {
		return nil, fmt.Errorf("global: %v", err)
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	data := struct {
		Meta struct {
			GeneratedAt string `json:"generated_at"`
		} `json:"meta"`
		Template ImageTemplate `json:"template"`
		// emotions are stored in a dynamic format, the key is the name of the emoticon
		Emotions map[string]struct {
			Description string `json:"description"`
			ImageID     int    `json:"image_id"`
			FirstSeen   string `json:"first_seen"`
		} `json:"emotes"`
	}{}
	if err := dec.Decode(&data); err != nil {
		return nil, fmt.Errorf("global: %v", err)
	}
	t, err := time.Parse(time.RFC3339, data.Meta.GeneratedAt)
	if err != nil {
		log.Printf("could not parse 'generated at' time %q: %v\n", data.Meta.GeneratedAt, err)
	}
	res := GlobalResult{
		Meta: Metadata{GeneratedAt: t},
		Template: ImageTemplate{
			Small:  data.Template.Small,
			Medium: data.Template.Medium,
			Large:  data.Template.Large,
		},
		Emotes: make([]Emoticon, 0, len(data.Emotions)),
	}
	for k, v := range data.Emotions {
		t := time.Time{}
		if v.FirstSeen != "" {
			t, err = time.Parse("2006-01-02 15:04:05", v.FirstSeen)
			if err != nil {
				log.Printf("could not parse time %q of emoticon %q: %v\n", v.FirstSeen, k, err)
			}
		}
		res.Emotes = append(res.Emotes, Emoticon{
			Code:        k,
			Description: v.Description,
			ImageID:     v.ImageID,
			FirstSeen:   t,
		})
	}
	return &res, nil
}
