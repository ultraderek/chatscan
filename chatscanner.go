package main

import (
	"chatscan/chatscan"
	"flag"
	"fmt"
)

func main() {

	jsonlocation := flag.String("chat", "nofile", "Chat File")
	newfilelocation := flag.String("csv", "output", "This is the CSV output file name .csv will be added itself")
	emotelimit := flag.Int("elimit", 5, "the number of emotes that you want to track if limit is less than one or to much total emotes will stats will print")
	excludebeginning := flag.Int("exbegins", 0, "the num of seconds excluded from beginning of file")
	excludeend := flag.Int("exends", 0, "the num of seconds excluded from end of file")
	steps := flag.Int("step", 30, "the size of steps when counting in seconds")
	overlap := flag.Bool("overlap", false, "this overlaps the steps making double counts")
	flag.Parse()
	size := *steps
	if *overlap {
		size = *steps * 2
	}
	fmt.Println(*jsonlocation, *newfilelocation, *emotelimit, *excludebeginning, *excludeend, *steps, *overlap)
	if *jsonlocation == "nofile" {
		fmt.Println("No file chosen use flag -h for help")
		return
	}
	chatscan.ProgramMain2(*newfilelocation, *jsonlocation, *emotelimit, *excludebeginning, *excludeend, *steps, size)

}
