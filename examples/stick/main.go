package main

import (
	"flag"
	"fmt"
	"github.com/hectormendez/bobblehat/sense/stick"
	"os"
	"os/signal"
)

func main() {
	var path string //this is the name of the vriable that allow read the output in direct

	flag.StringVar(&path, "path", "/dev/input/event0", "path to the event device") //this is the function that allows define which /dev/input/ do you want to reas

	// Parse command line flags
	flag.Parse()

	// Open the input device (and defer closing it)
	input, err := stick.Open(path)
	if err != nil {
		fmt.Printf("Unable to open input device: %s\nError: %v\n", path, err) //if the path is not correct the this alert that the endpoint do not run
		os.Exit(1)
	}

	// Print the name of the input device
	fmt.Println(input.Name()) /////This is not mandatory, but is usable because help you to idenify the devices :)

	// Set up a signals channel (stop the loop using Ctrl-C)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	// Loop forever capturing letters in upper case
	for {
		select {
		case <-signals:
			fmt.Println("")
			// Exit the loop
			return
		case e := <-input.Events:
			//fb := screen.NewFrameBuffer()
			//fmt.Printf(",%d,",e.Code)
			switch e.Code {
			case stick.Enter:
				fmt.Print("ENTER\n")
			case stick.Tab:
				fmt.Print("    TAB    ")
			case stick.Cero:
				fmt.Print("0")
			case stick.Uno:
				fmt.Print("1")
			case stick.Dos:
				fmt.Print("2")
			case stick.Tres:
				fmt.Print("3")
			case stick.Cuatro:
				fmt.Print("4")
			case stick.Cinco:
				fmt.Print("5")
			case stick.Seis:
				fmt.Print("6")
			case stick.Siete:
				fmt.Print("7")
			case stick.Nueve:
				fmt.Print("9")
			case stick.A:
				fmt.Print("a")
			case stick.B:
				fmt.Print("b")
			case stick.C:
				fmt.Print("c")
			case stick.D:
				fmt.Print("d")
			case stick.E:
				fmt.Print("e")
			case stick.F:
				fmt.Print("f")
			case stick.Te:
				fmt.Print("t")
			case stick.Dot:
				fmt.Print(":")

				//If do you want add more letter is posible, the dictionary is already complete

			}

		}
	}
}

/*func draw(fb *screen.FrameBuffer, a, b, m, n int, c color.Color) {
	for i := a; i < m; i++ {
		for j := b; j < n; j++ {
			fb.SetPixel(i, j, c)
		}
	}
}
*/
