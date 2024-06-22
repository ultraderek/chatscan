package chatscan

import (
	"chatscan/chatjson/twchat"
	"fmt"
	"os"
)

const newfilelocation = "sampleout"
const jsonlocation = "chat.json"

/*
const basedirectory = "/home/derek/Desktop/edits/edits todo"
const jsonlocation = basedirectory + "/Pippa/Pippa20240528HerLastStream/chat.json"
const newfilelocation = basedirectory + "/Pippa/Pippa20240528HerLastStream/chatdata"
const basedirectory = "/home/derek/Desktop/edits/edits todo"
const jsonlocation = basedirectory + "/henya/Henya20240428DarkSouls/chat.json"
const newfilelocation = basedirectory + "/henya/Henya20240428DarkSouls/chatdata"
*/
//const jsonlocation = "/home/derek/Desktop/edits/edits todo/LucyPyre/lucychat.json"

type overlapingseconds struct {
	starttime string
	startval  uint
	sumedvals int
}

// ProgramMain is the main function for this program
func ProgramMain() {

}

// ProgramMain2 is a test program
func ProgramMain2() {

	excludebeginning := 56 * 60
	//5:38:52-3:45:12 =1:53:40
	excludeend := 6820

	steps := 30
	size := steps * 2

	writefile, err := os.Create(fmt.Sprintf("%v%v.csv", newfilelocation, steps))
	if err != nil {
		panic(err)
	}
	defer writefile.Close()

	feed, err := twchat.CreateFeed(jsonlocation)
	if err != nil {
		panic(err)
	}

	cntr := make([]messagecounter, int(feed[len(feed)-1].TimeInSeconds)+1)
	for _, x := range feed {
		if x.TimeInSeconds >= 0 {
			cntr[int(x.TimeInSeconds)].loadenames(x)
			cntr[int(x.TimeInSeconds)].timetext = x.TimeText
		}

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

				section.sumedvals += cntr[j].nmessages
			}
		}
		sumedsecsarray = append(sumedsecsarray, section)
	}
	var sumation int
	for x := range sumedsecsarray {
		if x >= (excludebeginning/steps) && x <= len(sumedsecsarray)-(excludeend/steps) {
			sumation += sumedsecsarray[x].sumedvals
		}

	}
	sumation /= (len(sumedsecsarray) - ((excludeend + excludebeginning) / steps))

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
