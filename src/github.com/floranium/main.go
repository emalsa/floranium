package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	//f, err := os.Create("data_log.txt")

	f, err := os.OpenFile("/Users/daniele/projects/floranium/src/github.com/floranium/dataLog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	replaceCompiled := regexp.MustCompile(`\;{1,999}`)
	for input.Scan() {
		replacedString := input.Text()
		replacedString = strings.Replace(replacedString, " ", ";", -1)
		replacedString = strings.Replace(replacedString, "*", "", -1)
		replacedString = replaceCompiled.ReplaceAllString(replacedString, ";")

		firstCharacter := replacedString[0:1]
		lastCharacter := replacedString[len(replacedString)-1:]
		if firstCharacter == ";" {
			replacedString = RemoveFirstChar(replacedString)
		}
		if lastCharacter == ";" {
			replacedString = strings.TrimSuffix(replacedString, ";")
		}

		currentTime := time.Now()
		dateTime := currentTime.Format("2006-01-02 15:04:05.000000000")

		replacedString = replacedString + ";" + dateTime
		if _, err := f.Write([]byte(replacedString + "\n")); err != nil {
			log.Fatal(err)
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func RemoveFirstChar(input string) string {
	if len(input) <= 1 {
		return ""
	}
	return input[1:]
}

func RemoveLastChar(input string) string {
	if len(input) <= 1 {
		return ""
	}
	return input[1:]
}

// cd ~/projects/floranium/src/github.com/floranium && go build
// minicom -C raw.txt | tail -f raw.txt  | ./floranium
