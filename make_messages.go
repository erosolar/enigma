package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.mit.edu/erosolar/enigma/encoder"
	"github.mit.edu/erosolar/enigma/messages"
)

func main() {
	settings := encoder.Settings{}
	settings.RotorOrder = []int{5,1,2}
	settings.RingSettings = []int{14,4,12}
	settings.Plugs = "AN IV LH YP WM TR XU FO ZB ED"
	settings.Reflector = encoder.UKWB

	files, err := ioutil.ReadDir("messages/plaintext")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		inputFileName := "messages/plaintext/" + file.Name()
		outputFileName := "messages/ciphertext/" + file.Name()

		if err := messages.Run(inputFileName, outputFileName, settings); err != nil {
			fmt.Println("Error with message creation in file", file.Name(), "--", err)
		} else {
			fmt.Println("Successfully encrypted today's messages for", file.Name())
		}
	}
}

func interactiveMessageMaking() {

	fmt.Printf("Welcome to Enigma message creation.\n")
	settings := encoder.Settings{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter today's rotor settings (eg. 1 2 3): ")
	text, _ := reader.ReadString('\n')
	settings.RotorOrder = make([]int, 0, 3)
	for _, num := range strings.Split(text[0:len(text)-1], " ") {
		if num != "" {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			settings.RotorOrder = append(settings.RotorOrder, i)
		}
	}

	fmt.Print("Enter today's ring settings (eg. 01 24 03): ")
	text, _ = reader.ReadString('\n')
	settings.RingSettings = make([]int, 0, 3)
	for _, num := range strings.Split(text[0:len(text)-1], " ") {
		if num != "" {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			settings.RingSettings = append(settings.RingSettings, i)
		}
	}

	fmt.Print("Enter today's plug pairs (eg. GP XH TW IA ...): ")
	text, _ = reader.ReadString('\n')
	settings.Plugs = text[0 : len(text)-1]

	settings.Reflector = encoder.UKWB

	fmt.Print("Enter the file name for today's messages: ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	inputFileName := "messages/plaintext/" + fileName
	outputFileName := "messages/ciphertext/" + fileName

	if err := messages.Run(inputFileName, outputFileName, settings); err != nil {
		fmt.Println("Error with message creation:", err)
	} else {
		fmt.Println("Successfully encrypted today's messages")
	}
}
