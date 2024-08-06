package chatscan

import (
	"strings"
)

type emotecounter struct {
	emote string
	n     int
}
type messagecounter struct {
	emotes    []emotecounter
	nmessages int
	timetext  string
	//time      int
}
type EmoteNamer interface {
	EmoteNames() (s []string)
}

func (m *messagecounter) append(c *messagecounter) {
	m.nmessages += c.nmessages
	if c.emotes == nil {
		return
	}
	for i, emote := range c.emotes {
		hasemote := false
		for j, x := range m.emotes {
			if strings.Contains(x.emote, emote.emote) {
				m.emotes[j].n += emote.n
				hasemote = true
				break
			}
		}
		if !hasemote {
			m.emotes = append(m.emotes, c.emotes[i])
		}
	}
}
func (m *messagecounter) loadenames(x EmoteNamer) {
	m.nmessages++
	emotearray := x.EmoteNames()

	if emotearray == nil {
		return
	}

	for _, emote := range emotearray {
		hasemote := false
		for i, x := range m.emotes {
			if strings.Contains(x.emote, emote) {

				m.emotes[i].n++
				hasemote = true
				break
			}
		}
		if !hasemote {
			m.emotes = append(m.emotes, emotecounter{emote: emote, n: 1})
		}
	}

}

// totalmessages returns the number of messages that were counted in sloted time
func (m *messagecounter) totalmessages() int {
	return m.nmessages
}
