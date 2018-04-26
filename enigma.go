package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.mit.edu/erosolar/enigma/encoder"
)

var enigma encoder.Enigma
var reader *bufio.Reader

func main() {
	fmt.Printf("Welcome to Enigma.\n")
	settings := encoder.Settings{}

	reader = bufio.NewReader(os.Stdin)
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

	enigma = encoder.Setup(settings)

	for {
		fmt.Println("ENCODE or DECODE?")
		text, _ = reader.ReadString('\n')
		if text == "EXIT\n" {
			return
		} else if text[0] == 'E' {
			encrypt()
		} else if text[0] == 'D' {
			decrypt()
		} else {
			fmt.Println("enter \"ENCODE\" or \"DECODE\" please!")
		}
	}
}

func encrypt() {
	fmt.Print("enter outer key: ")
	key, _ := reader.ReadString('\n')
	enigma = encoder.Initialize(enigma, key[0:3])
	fmt.Print("enter inner key: ")
	key2, _ := reader.ReadString('\n')
	output := enigma.Encrypt(key2[0:3] + key2[0:3])

	enigma = encoder.Initialize(enigma, key2[0:3])
	fmt.Println("enter message below")
	text, _ := reader.ReadString('\n')
	fmt.Println("encrypted message:")
	output += enigma.Encrypt(text[0 : len(text)-1])
	fmt.Println(output)
}

func decrypt() {
	fmt.Print("enter (outer) key: ")
	key, _ := reader.ReadString('\n')
	enigma = encoder.Initialize(enigma, key[0:3])
	fmt.Println("enter ciphertext below")
	text, _ := reader.ReadString('\n')
	if len(text) < 6 {
		fmt.Println("ciphertext must be at least 6ch long")
		return
	}
	key2 := enigma.Encrypt(text[0:3])
	enigma = encoder.Initialize(enigma, key2)
	fmt.Println(enigma.Encrypt(text[6 : len(text)-1]))

}
