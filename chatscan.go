package main

import (
	"encoding/json"
	"io"

	"fmt"
	"os"
)

const basedirectory = "/home/derek/Desktop/edits/edits todo"
const jsonlocation = basedirectory + "/henya/Henya20240428DarkSouls/chat.json"
const newfilelocation = basedirectory + "henya/Henya20240428DarkSouls/chatdata"

//const jsonlocation = "/home/derek/Desktop/edits/edits todo/LucyPyre/lucychat.json"

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
	Version     uint   `json:"version"`
}
type Emote struct {
	// Images []Image `json:"images"`
	ID        string `json:"id"`
	Locations string `json:"locations"`
	Name      string `json:"name"`
}

func CreateTwitchFeed() (x *TwitchFeed) {
	return &TwitchFeed{}
}

type overlapingseconds struct {
	starttime string
	startval  uint
	sumedvals int
}

func main() {
	steps := 15
	size := steps * 2
	file, err := os.Open(jsonlocation)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writefile, err := os.Create(fmt.Sprintf("%v%v", newfilelocation, steps))
	if err != nil {
		panic(err)
	}
	defer writefile.Close()
	//bufreader := bufio.NewReader(file)
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	twfeed := make([]Message, 0)
	err = json.Unmarshal(data, &twfeed)
	if err != nil {
		panic(err)
	}

	cntr := make([]int, twfeed[len(twfeed)-1].TimeSeconds+1)
	for _, x := range twfeed {
		cntr[x.TimeSeconds]++
	}

	sumedsecsarray := make([]overlapingseconds, 0)
	for i := 0; i < len(cntr); i = i + steps {
		minutes := i / 60
		hours := minutes / 60
		seconds := i % 60
		minutes = minutes % 60

		section := overlapingseconds{startval: uint(i), starttime: fmt.Sprintf("%v:%v:%v", hours, minutes, seconds)}
		for j := i; j < i+size; j++ {
			if j < len(cntr) {
				section.sumedvals += cntr[j]
			}
		}
		sumedsecsarray = append(sumedsecsarray, section)
	}
	var sumation int
	for x := range sumedsecsarray {
		sumation += sumedsecsarray[x].sumedvals
	}
	sumation /= (len(sumedsecsarray))

	for x := range sumedsecsarray {
		if sumedsecsarray[x].sumedvals-sumation > 0 {
			_, err = fmt.Fprintf(writefile, "%v,%v,%v,%v\n", sumedsecsarray[x].startval,
				sumedsecsarray[x].starttime,
				sumedsecsarray[x].sumedvals,
				sumedsecsarray[x].sumedvals-sumation)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			_, err = fmt.Fprintf(writefile, "%v,%v,%v,%v\n", sumedsecsarray[x].startval,
				sumedsecsarray[x].starttime,
				sumedsecsarray[x].sumedvals,
				0)
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
