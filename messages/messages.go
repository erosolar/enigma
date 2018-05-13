package messages

import (
	"bufio"
	"math/rand"
	"os"
	"strings"

	"github.mit.edu/erosolar/enigma/encoder"
)

const (
	alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	comma            = "ZZ"
	questionMark     = "FRAQ"
	period           = "X"
    ampersand        = "AND"
    exclamationPoint = "YUD"
	// numbers
	zero  = "zero"
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

var enigma encoder.Enigma

func Run(inputFileName, outputFileName string, settings encoder.Settings) error {
	messages, err := readFile(inputFileName)
	for err != nil {
		return err
	}

	return encryptAndWriteToFile(messages, outputFileName, settings)
}

func readFile(fileName string) ([]string, error) {
	messages := []string{}

	file, err := os.Open(fileName)
	if err != nil {
		return messages, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	startEncrypting := false
	for scanner.Scan() {
		text := scanner.Text()
		// the text of the book comes after the line with the *** in it
		if strings.Contains(text, "***") {
			startEncrypting = true
			continue
		}
		if startEncrypting && text != "" {
			messages = append(messages, makeMessage(text))
		}
	}

	if err = scanner.Err(); err != nil {
		return messages, err
	}

	return messages, nil
}

func makeMessage(input string) string {
	return formatMessage(strings.Fields(input))
}

func formatMessage(s []string) string {
	r := strings.NewReplacer(",", comma, "?", questionMark, ".", period, "!", exclamationPoint, "&", ampersand,
        "0", zero, "1", one, "2", two, "3", three, "4", four, "5", five, "6", six, "7", seven, "8", eight, "9", nine)

	result := strings.ToUpper(r.Replace(strings.Join(s, "")))

	f := func(c rune) bool { return c < 65 || c > 90 }

	result = strings.Join(strings.FieldsFunc(result, f), "")

	// pad to be group of 5
	for i := 0; i < len(result)%5; i++ {
		result += "X"
	}

	return result
}

func encryptAndWriteToFile(messages []string, fileName string, settings encoder.Settings) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// just so it's different
	rand.Seed(int64(settings.RotorOrder[0] + settings.RingSettings[2]))
	enigma = encoder.Setup(settings)

	w := bufio.NewWriter(file)
	for _, m := range messages {
		_, err = w.WriteString(encryptMessage(m) + "\n")
		if err != nil {
			return err
		}
	}
	w.Flush()

	return nil
}

func encryptMessage(s string) string {
	outerKey := randomKey()
	innerKey := randomKey()

	enigma = encoder.Initialize(enigma, outerKey)
	output := enigma.Encrypt(innerKey)
	enigma = encoder.Initialize(enigma, innerKey)
	output += enigma.Encrypt(s + "XX") // to account for the 3 extra characters in the inner key

	// format for "transmission"
	encryptedMsg := ""
	var i int
	for i = 0; i < len(output); i += 5 {
		encryptedMsg += output[i:i+5] + " "
	}

	return outerKey + " " + encryptedMsg[0:len(encryptedMsg)]
}

func randomKey() string {
	return string(alpha[rand.Intn(26)]) + string(alpha[rand.Intn(26)]) + string(alpha[rand.Intn(26)])
}
