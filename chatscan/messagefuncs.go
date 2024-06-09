package chatscan

import "strings"

type emotecounter struct {
	emote string
	n     int
}
type messagecounter struct {
	emotes    []emotecounter
	nmessages int
	timetext  string
}
type EmoteNamer interface {
	EmoteNames() (s []string)
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
