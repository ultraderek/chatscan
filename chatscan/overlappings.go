package chatscan

import (
	"errors"
	"fmt"
	"sort"
)

type overlapingseconds struct {
	starttime string
	startval  uint
	sumedvals int
	emotes    []sumedemotes
	csvemotes []sumedemotes
}
type sumedemotes struct {
	name      string
	sumedvals int
}
type arrayofsumedemotes []sumedemotes

func (a arrayofsumedemotes) Len() int           { return len(a) }
func (a arrayofsumedemotes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a arrayofsumedemotes) Less(i, j int) bool { return a[i].sumedvals > a[j].sumedvals }

// csvstring - this output a format with for csv which includes
// 1)time in seconds (s)
// 2)time in text (hh:mm:ss)
// 3)not ajusted message count (count)
// 4)adjusted message count (formula is (count - average count) >0)
// 5)
// ...)Emotecount (counts the emotes used)
// N)
func (m *overlapingseconds) csvstring(adjustoffsetvalue int) (value string, err error) {

	if m.csvemotes == nil {
		return value, errors.New("csvstring - inner memory is nil try running csvprepwork")
	}
	if adjustoffsetvalue < m.sumedvals {
		value = fmt.Sprintf("%v,%v,%v,%v", m.startval,
			m.starttime,
			m.sumedvals,
			m.sumedvals-adjustoffsetvalue)
	} else {
		value = fmt.Sprintf("%v,%v,%v,%v", m.startval,
			m.starttime,
			m.sumedvals,
			0)
	}

	for i := 0; i < len(m.csvemotes); i++ {
		value = value + "," + fmt.Sprint(m.csvemotes[i].sumedvals)
	}

	return value, nil
}
func (m *overlapingseconds) csvprepwork(tots []sumedemotes) {
	m.csvemotes = make([]sumedemotes, 0)
	for _, metaemote := range tots {
		hasemote := false
		for _, localemote := range m.emotes {
			if localemote.name == metaemote.name {
				hasemote = true
				m.csvemotes = append(m.csvemotes, localemote)
				break
			}

		}
		if !hasemote {
			m.csvemotes = append(m.csvemotes, sumedemotes{metaemote.name, 0})
		}
	}
}

// findcsvheader if numofemotes is negative do all
func createcsvheader(tots []sumedemotes) (header string) {
	header = fmt.Sprintf("%v,%v,%v,%v", "TOS", "TextTime", "Count", "AdjCount")

	for i := 0; i < len(tots); i++ {
		header = header + "," + tots[i].name
	}
	return header

}

func emotetotals(data []overlapingseconds, numofemotes int) (tots []sumedemotes) {
	tots = make([]sumedemotes, 0)
	foundemotes := make([]sumedemotes, 0)
	for _, section := range data {
		for _, emotesums := range section.emotes {
			emote := emotesums
			contains := false
			for i := 0; i < len(foundemotes); i++ {

				//if strings.Contains(emote.name, foundemotes[i].name) {
				if emote.name == foundemotes[i].name {
					contains = true
					foundemotes[i].sumedvals += emote.sumedvals
					break
				}
			}
			if !contains {
				foundemotes = append(foundemotes, emote)
			}
		}
	}
	tots = append(tots, foundemotes...)
	sort.Sort(arrayofsumedemotes(tots))
	if numofemotes > len(tots) || numofemotes < 1 {
		return tots
	}

	return tots[:numofemotes]
}
