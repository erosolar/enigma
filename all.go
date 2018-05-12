package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.mit.edu/erosolar/enigma/bombe"
	"github.mit.edu/erosolar/enigma/checker"
	"github.mit.edu/erosolar/enigma/menumaker"
)

type Message struct {
	plaintextKey string
	encryptedKey string
	message      string
}

func main() {
	fmt.Println("Welcome to the Engima decoding program.")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the file name for today's messages: ")
	messageFileName, _ := reader.ReadString('\n')
	messageFileName = strings.TrimSpace(messageFileName)
	
	messages, err := readMessageFile(messageFileName)
	for err != nil {
		fmt.Print("Error reading file ", messageFileName, ", please enter it in again: ")
		messageFileName, _ := reader.ReadString('\n')
		messages, err = readMessageFile(messageFileName)
	}

	fmt.Print("Enter the message crib: ")
	crib, _ := reader.ReadString('\n')

	fmt.Println("Please wait while we work.")

	killChan := make(chan bool)
	menuChan := make(chan menumaker.Menu)
	resultChan := make(chan bombe.Result)
	userChan := make(chan bombe.Result)

	numBombes := 10

	go makeMenus(messages, crib, menuChan, killChan)
	go startBombes(numBombes, menuChan, resultChan, killChan)
	go runChecker(resultChan, userChan, killChan)

	timer := time.After(time.Duration(rand.Intn(10)) * time.Second)

	for {
		select {
		case res := <-userChan:
			fmt.Printf("\nPossible settings -- offset: %v; rotors: %v\n", res.Offset, res.Rotors)
			fmt.Println(res.Printable)
			// TODO interact with user to figure out turnover/if it's a good setting
			// if we found today's settings
			fmt.Print("Exit? (y/n) ")
			exit, _ := reader.ReadString('\n')
			if exit == "y" {
				close(killChan)
				return
			}
		case <-timer:
			notes := []string{"\u2669", "\u266A", "\u266B", "\u266C"}
			fmt.Print(notes[rand.Intn(4)], " ")
			timer = time.After(time.Duration(rand.Intn(10)) * time.Second)
		}
	}
}

func readMessageFile(fileName string) ([]Message, error) {
	messages := []Message{}

	file, err := os.Open(fileName)
	if err != nil {
		return messages, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		messages = append(messages, makeMessage(scanner.Text()))
	}

	if err = scanner.Err(); err != nil {
		return messages, err
	}

	return messages, nil
}

func makeMessage(input string) Message {
	return Message{
		plaintextKey: input[0:3],
		encryptedKey: input[3:9],
		message:      input[9:len(input)],
	}
}

func makeMenus(messages []Message, crib string, menuChan chan menumaker.Menu, killChan chan bool) {
	for _, msg := range messages {
		menus := menumaker.MakeMenus(msg.message, crib)
		for _, m := range menus {
			go func(menu menumaker.Menu, menuChan chan menumaker.Menu, killChan chan bool) {
				select {
				case menuChan <- menu:
				case <-killChan:
				}
			}(m, menuChan, killChan)
		}
	}
}

func startBombes(numBombes int, menuChan chan menumaker.Menu, resultChan chan bombe.Result, killChan chan bool) {
	for i := 0; i < numBombes; i++ {
		go startBombe(menuChan, resultChan, killChan)
	}
}

func startBombe(menuChan chan menumaker.Menu, resultChan chan bombe.Result, killChan chan bool) {
	for {
		select {
		case m := <-menuChan:
		     runBombe(m, resultChan, killChan)
		case <-killChan:
			return
		}
	}
}

func runBombe(menu menumaker.Menu, resultChan chan bombe.Result, killChan chan bool) {
     // TODO
}

func runChecker(resultChan chan bombe.Result, userChan chan bombe.Result, killChan chan bool) {
	for {
		select {
		case r := <-resultChan:
			if checker.CheckIfPossiblePlugboard(r.State) {
				go func(r bombe.Result, userChan chan bombe.Result, killChan chan bool) {
					select {
					case userChan <- r:
					case <-killChan:
					}
				}(r, userChan, killChan)
			}
		case <-killChan:
			return
		}
	}
}
