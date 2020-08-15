package main

import (
	"bufio"
	"fmt"
	"musicbox/musicpack"
	"os"
)

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	exit := false
	fmt.Println("Type in mp3 track you wish to play: ")
	Scanner.Scan()
	result2 := Scanner.Text()
	go musicpack.Play(result2 + ".mp3")
	for exit == false {
		fmt.Println("Press 'h' for additional instruction: ")
		Scanner.Scan()
		result := Scanner.Text()
		switch result {
		case "h":
			fmt.Println(`
			q - quit
			w - speed music up
			s - slow music down
			l - load a new track
			k - stop music
			'-' - turn it down
			'+' - turn it up
			`)
		case "q":
			exit = true
		case "w":
			go musicpack.Up()
		case "s":
			go musicpack.Down()
		case "k":
			go musicpack.Stop()
		case "-":
			go musicpack.Quiet()
		case "+":
			go musicpack.Loud()
		case "l":
			go musicpack.Stop()
			fmt.Println("Type in new track name: ")
			Scanner.Scan()
			result2 := Scanner.Text()
			go musicpack.Play(result2 + ".mp3")
		}
	}
}
