package ytchat

import (
	"encoding/json"
	"io"
	"os"
)

type Message struct {
	ActionType    string  `json:"action_type"`
	Author        Author  `json:"author"`
	Emotes        []Emote `json:"emotes"`
	Message       string  `json:"message"`
	MessageType   string  `json:"message_type"`
	TimeInSeconds float64 `json:"time_in_seconds"`
	TimeText      string  `json:"time_text"`
	Timestamp     int     `json:"timestamp"`
}
type Author struct {
	Badges []Badge `json:"badges"`
	ID     string  `json:"id"`
	Images []Image `json:"images"`
	Name   string  `json:"name"`
}
type Emote struct {
	Images        []Image `json:"images"`
	IsCustomEmoji bool    `json:"is_custom_emoji"`
	Name          string  `json:"name"`
}
type Badge struct {
	Icons []Icon `json:"icons"`
	Title string `json:"title"`
}

type Image struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
type Icon struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (m *Message) GetAuthorName() string {
	return m.Author.Name
}

func (m Message) EmoteNames() (s []string) {
	for _, x := range m.Emotes {
		s = append(s, x.Name)
	}
	return s
}
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
