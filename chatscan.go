package main

import (
	"bufio"
	"chatscans/chatjson"
	"fmt"
	"os"
	"strings"
)

func printingstrings(x *bufio.Reader) error {
	d := []byte("\n")
	i := 0
	for {
		str, err := x.ReadString(d[0])

		fmt.Println(i, str)
		if err != nil {
			return err
		}
		i++
	}

}

type emotecounter struct {
	emote   string
	counter uint
}
type timescounter struct {
	hhmmss  string
	counter uint
	//nemotes []emotecounter
}
type timesarray struct {
	x   []timescounter
	pos uint
}

func (t *timesarray) append(time string) {
	if strings.Contains(t.x[t.pos].hhmmss, time) {
		t.x[t.pos].counter++
	}
	t.pos++
	t.x = append(t.x, timescounter{time, 1})

}

func main() {
	twfeed := chatjson.CreateTwitchFeed()

	file, err := os.Open("chatlog")
	if err != nil {
		panic(err)
	}
	bufreader := bufio.NewReader(file)

	fmt.Print(printingstrings(bufreader))
	fmt.Println(twfeed)
}
