package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.mit.edu/erosolar/enigma/bombe"
	"github.mit.edu/erosolar/enigma/checker"
	"github.mit.edu/erosolar/enigma/encoder"
	"github.mit.edu/erosolar/enigma/menumaker"
)

type message struct {
	plaintextKey string
	encryptedKey string
	message      string
}

type bombeMessage struct {
	message string
	menu    menumaker.Menu
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
	crib = strings.TrimSpace(crib)

	fmt.Println("Please wait while we work.")

	killChan := make(chan bool)
	menuChan := make(chan bombeMessage)
	resultChan := make(chan bombe.Result)
	userChan := make(chan bombe.Result)

	numBombes := 10

	go makeMenus(messages, crib, menuChan, killChan)
	go startBombes(numBombes, menuChan, resultChan, killChan)
	go runChecker(resultChan, userChan, killChan)

	results := []bombe.Result{}

L:
	for {
		select {
		case res, ok := <-userChan:
			if !ok {
				break L
			}
			notes := []string{"\u2669", "\u266A", "\u266B", "\u266C"}
			fmt.Print(notes[rand.Intn(4)], " ")
			results = append(results, res)
		}
	}

	// post-process possible settings
	fmt.Println("Congratulations, you've made it post-processing!")

	fmt.Println("Number of results:", len(results))

	processResults(results)
}

func readMessageFile(fileName string) ([]message, error) {
	messages := []message{}

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

func makeMessage(input string) message {
	return message{
		plaintextKey: input[0:3],
		encryptedKey: input[3:6],
		message:      strings.Replace(input[6:len(input)], " ", "", -1),
	}
}

func makeMenus(messages []message, crib string, menuChan chan bombeMessage, killChan chan bool) {
	for _, msg := range messages {
		menus := menumaker.MakeMenus(msg.message, crib)
		for _, m := range menus {
			if m.NumLetters > 10 {
				select {
				case menuChan <- bombeMessage{msg.message, m}:
				case <-killChan:
				}
			}
		}
	}
	close(menuChan)
	fmt.Print("\u22EE")
}

func startBombes(numBombes int, menuChan chan bombeMessage, resultChan chan bombe.Result, killChan chan bool) {
	doneChan := make(chan bool)
	for i := 0; i < numBombes; i++ {
		go startBombe(menuChan, resultChan, doneChan, killChan)
	}
	for i := 0; i < numBombes; i++ {
		<-doneChan
	}
	close(resultChan)
	fmt.Print("\u22EE")
}

func startBombe(menuChan chan bombeMessage, resultChan chan bombe.Result, doneChan chan bool, killChan chan bool) {
	for {
		select {
		case bm, ok := <-menuChan:
			if !ok {
				doneChan <- true
				return
			}
			runBombe(bm, resultChan, killChan)
		case <-killChan:
			return
		}
	}
}

func runBombe(bombeMessage bombeMessage, resultChan chan bombe.Result, killChan chan bool) {
	settings := bombe.Settings{
		Connections: bombeMessage.menu.Connections,
		NumEnigmas:  len(bombeMessage.menu.Connections),
		NumLetters:  bombeMessage.menu.NumLetters,
	}
	bombe.GetResults(bombeMessage.message, settings, 3, resultChan, killChan)
}

func runChecker(resultChan chan bombe.Result, userChan chan bombe.Result, killChan chan bool) {
	for {
		select {
		case r, ok := <-resultChan:
			if !ok {
				close(userChan)
				fmt.Print("\u22EE")
				return
			}
			if pb, ok := checker.CheckIfPossiblePlugboard(r.State); ok {
				r.Plugboard = pb
				select {
				case userChan <- r:
				case <-killChan:
				}
			}
		case <-killChan:
			return
		}
	}
}

func makeKey(offset int) string {
	high := offset / (26 * 26)
	med := (offset / 26) % 26
	low := offset % 26

	return string([]rune{rune(high + 'A'), rune(med + 'A'), rune(low + 'A')})
}

func processResults(results []bombe.Result) {
	sorted := make(map[string][]bombe.Result)

	for _, res := range results {
		if _, ok := sorted[stringify(res.Rotors)]; ok {
			sorted[stringify(res.Rotors)] = append(sorted[stringify(res.Rotors)], res)
		} else {
			sorted[stringify(res.Rotors)] = []bombe.Result{res}
		}
	}

	reader := bufio.NewReader(os.Stdin)
	for rotors, res := range sorted {
		steckerOpts := combineSteckers(res)
		for _, st := range steckerOpts {
			settings := encoder.Settings{
				RotorOrder:   []int{res[0].Rotors[2], res[0].Rotors[1], res[0].Rotors[0]}, // reversed >.>
				RingSettings: []int{1, 1, 1},                                              //assume ignored
				Plugs:        st,
				Reflector:    encoder.UKWB,
			}
			enig := encoder.Setup(settings)
			fmt.Printf("option: %s rotors with %s steckers\n", rotors, st)
			for _, r := range res {
				e := encoder.Initialize(enig, makeKey(r.Offset))
				fmt.Println(e.Encrypt(r.Message))
			}
			fmt.Print("is this a possible setting? (y/N)>")
			resp, _ := reader.ReadString('\n')
			if resp == "n\n" || resp == "N\n" || resp == "\n" {
				continue
			} else {
				fmt.Println("where does it kind of stop making sense?")
				resp, _ = reader.ReadString('\n')
				fmt.Println(resp)
			}
		}
	}

}

// expects list of results with same rotor order
func combineSteckers(results []bombe.Result) []string {
	out := make([]string, 0) // stecker lists
	for _, r := range results {
		inserted := false
		steckers := r.Plugboard
		for i, l := range out {
			if comb, ok := validCombination(steckers, l); ok {
				inserted = true
				out[i] = comb
			}
		}
		if !inserted {
			out = append(out, steckers)
		}
	}
	return out
}

func validCombination(l1, l2 string) (string, bool) {
	state := make(map[int]map[int]bool, len(l1))
	for _, pair := range strings.Split(l1, " ") { // add l1 pairs
		if _, ok := state[int(pair[0]-'A')]; !ok {
			state[int(pair[0]-'A')] = make(map[int]bool, 1)
		}
		state[int(pair[0]-'A')][int(pair[1]-'A')] = true
	}
	for _, pair := range strings.Split(l2, " ") { // add l2 pairs
		if _, ok := state[int(pair[0]-'A')]; !ok {
			state[int(pair[0]-'A')] = make(map[int]bool, 1)
		}
		state[int(pair[0]-'A')][int(pair[1]-'A')] = true
	}
	return checker.CheckIfPossiblePlugboard(state)
}

func stringify(l []int) string {
	s := ""
	for _, n := range l {
		s += strconv.Itoa(n)
	}
	return s
}
