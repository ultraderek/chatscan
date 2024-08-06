package chatscan

import (
	"chatscan/chatjson/twchat"
	"fmt"
	"os"
)

/*
const basedirectory = "/home/derek/Desktop/edits/edits todo"
const jsonlocation = basedirectory + "/Pippa/Pippa20240528HerLastStream/chat.json"
const newfilelocation = basedirectory + "/Pippa/Pippa20240528HerLastStream/chatdata"
const basedirectory = "/home/derek/Desktop/edits/edits todo"
const jsonlocation = basedirectory + "/henya/Henya20240428DarkSouls/chat.json"
const newfilelocation = basedirectory + "/henya/Henya20240428DarkSouls/chatdata"
*/
//const jsonlocation = "/home/derek/Desktop/edits/edits todo/LucyPyre/lucychat.json"

// ProgramMain is the main function for this program
func ProgramMain() {

}

// ProgramMain2 is a test program
func ProgramMain2(newfilelocation, jsonlocation string, emotelimit, excludebeginning, excludeend, steps, size int) {

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
			//fmt.Println(x)
			cntr[int(x.TimeInSeconds)].loadenames(x) //loadenamesalsosumsthemessages
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
				section.sumedvals += cntr[j].totalmessages()
				if len(cntr[j].emotes) != 0 {
					for _, x := range cntr[j].emotes {
						hasemote := false
						for k := 0; k < len(section.emotes); k++ {
							if section.emotes[k].name == x.emote {
								hasemote = true
								section.emotes[k].sumedvals += x.n
							}
						}
						if !hasemote {
							section.emotes = append(section.emotes, sumedemotes{x.emote, x.n})
						}

					}
				}
			}
		}
		sumedsecsarray = append(sumedsecsarray, section)
	}
	tots := emotetotals(sumedsecsarray, emotelimit)
	for _, section := range tots {
		fmt.Println(section)
	}
	var sumation int
	for x := range sumedsecsarray {
		if x >= (excludebeginning/steps) && x <= len(sumedsecsarray)-(excludeend/steps) {
			sumation += sumedsecsarray[x].sumedvals
		}
	}
	sumation /= (len(sumedsecsarray) - ((excludeend + excludebeginning) / steps))

	for i := 0; i < len(sumedsecsarray); i++ {
		sumedsecsarray[i].csvprepwork(tots)
	}
	_, err = fmt.Fprintf(writefile, "%v\n", createcsvheader(tots))
	if err != nil {
		fmt.Println(err)
	}
	for x := range sumedsecsarray {

		csvlineitem, err := sumedsecsarray[x].csvstring(sumation)
		if err != nil {
			fmt.Println(err)
		}
		_, err = fmt.Fprintf(writefile, "%v\n", csvlineitem)
		if err != nil {
			fmt.Println(err)
		}

	}

}
