package chatjson

type TwitchFeed struct {
	Messages []Message
}

type Message struct {
	Author      Author `json:"author"`
	Messsage    string `json:"message"`
	MessageType string `json:"message_type"`
	TimeSeconds uint   `json:"time_in_seconds"`
	TimeText    string `json:"time_text"`
	Timestamp   uint   `json:"timestamp"`
}

type Author struct {
	Badges      []Badge `json:"badges"`
	Emotes      []Emote `json:"emotes"`
	Color       string  `json:"colour"`
	DisplayName string  `json:"display_name"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
}

type Badge struct {
	ClickAction string `json:"clickAction"`
	ClickURL    string `json:"clickURL"`
	Name        string `json:"name"`
	Title       string `json:"title"`
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

func CreateTwitchFeed() (x *TwitchFeed) {
	return &TwitchFeed{}
}

// Get All Emotes returns all Emotes in a Message
// If no emotes s will have length of zero obviously
func (m Message) GetAllEmotes() (s []string) {
	for _, x := range m.Author.Emotes {
		s = append(s, x.Name)
	}
	return s
}
func (m Message) GetAllEmoteLinks() (s []string) {
	for _, x := range m.Author.Emotes {
		s = append(s, x.Images[len(x.Images)-1].URL)
	}
	return s
}
func (m Message) GetChatUserName() (s string) {
	return m.Author.Name
}
