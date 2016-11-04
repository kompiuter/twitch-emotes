package twitchemote

import "time"

// GlobalResult is the result of calling the Global emoticon API,
// containing all global Twitch emoticons
type GlobalResult struct {
	Meta     Metadata      `json:"metadata"`
	Template ImageTemplate `json:"template"`
	Emotes   []Emoticon    `json:"emotes"`
}

// Emoticon describes a Twitch emoticon. Code and ImageID are never empty, whereas
// the Description and FirstSeen fields may be empty (zero-value)
type Emoticon struct {
	// Code is the code to produce the emoticon in Twitch chat (i.e: Kappa)
	Code string `json:"code"`
	// ImageID is the id of the emoticon which may be combined with the Template from the metadata
	// to retrieve the image
	ImageID int `json:"image_id"`
	// Description describes the emoticon. Only applicable for global emoticons.
	Description string `json:"description"`
	// FirstSeen is the date where this emoticon was first introduced. Only applicable for global emoticons
	FirstSeen time.Time `json:"first_seen"`
}

// SubscriberResult is the result of calling the Subscriber emoticon API,
// containing all Twitch channels and their associated emoticons
type SubscriberResult struct {
	Meta     Metadata      `json:"metadata"`
	Template ImageTemplate `json:"template"`
	Channels []Channel     `json:"channels"`
}

// Channel describes a single Twitch channel and all associated information
// with it, including all emoticons for that channel
type Channel struct {
	Badge         string     `json:"badge"`
	Badge12m      string     `json:"badge_12m"`
	Badge24m      string     `json:"badge_24m"`
	Badge3m       string     `json:"badge_3m"`
	Badge6m       string     `json:"badge_6m"`
	BadgeStarting string     `json:"badge_starting"`
	ChannelID     string     `json:"channel_id"`
	Desc          string     `json:"desc"`
	Emotes        []Emoticon `json:"emotes"`
	FirstSeen     time.Time  `json:"first_seen"`
	ID            string     `json:"id"`
	Link          string     `json:"link"`
	Set           int        `json:"set"`
	Title         string     `json:"title"`
}

// Metadata describes the metadata retrieved from the global and subscriber APIs
type Metadata struct {
	GeneratedAt time.Time `json:"generated_at"`
}

// ImageTemplate describes the image URL for retrieving an emoticon's image
type ImageTemplate struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}
