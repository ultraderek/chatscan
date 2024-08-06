package twchat

import (
	"encoding/json"
	"io"
	"os"
)

type Message struct {
	Author        Author  `json:"author"`
	Emotes        []Emote `json:"emotes"`
	Messsage      string  `json:"message"`
	MessageType   string  `json:"message_type"`
	TimeInSeconds int     `json:"time_in_seconds"`
	TimeText      string  `json:"time_text"`
	Timestamp     uint    `json:"timestamp"`
}

type Author struct {
	Badges      []Badge `json:"badges"`
	Color       string  `json:"colour"`
	DisplayName string  `json:"display_name"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
}

type Badge struct {
	ClickAction string `json:"clickAction"`
	ClickURL    string `json:"clickURL"`
	Icons       []Icon `json:"icons"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	//Version     int    `json:"version"`
}
type Icon struct {
	Height int    `json:"height"`
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type Emote struct {
	ID        string  `json:"id"`
	Images    []Image `json:"images"`
	Locations string  `json:"locations"`
	Name      string  `json:"name"`
}
type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// CreateFeed creates an array of messages for twitch chat
func CreateFeed(path string) (m []Message, err error) {
	m = make([]Message, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Get All Emotes returns all Emotes in a Message
// If no emotes s will have length of zero obviously
func (m Message) EmoteNames() (s []string) {
	for _, x := range m.Emotes {
		s = append(s, x.Name)
	}
	return s
}
func (m Message) GetAllEmoteLinks() (s []string) {
	for _, x := range m.Emotes {
		s = append(s, x.Images[len(x.Images)-1].URL)
	}
	return s
}
func (m Message) GetChatUserName() (s string) {
	return m.Author.Name
}
