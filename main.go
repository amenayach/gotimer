package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func padLeft(str, pad string, length int) string {
	for {
		if len(str) >= length {
			return str
		}
		str = pad + str
		if len(str) > length {
			return str[0:length]
		}
	}
}

func displayDuration(seconds int) {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60
	fmt.Print("\r" +
		padLeft(strconv.Itoa(hours), "0", 2) + ":" +
		padLeft(strconv.Itoa(minutes), "0", 2) + ":" +
		padLeft(strconv.Itoa(secs), "0", 2) + "     ")
}

func getArgsSeconds(arg string) (seconds int, err error) {

	if len(arg) < 2 {
		return 0, errors.New("Invalid input")
	}

	timeType := arg[len(arg)-1:]
	timeStringValue := arg[0 : len(arg)-1]
	timeValue, err := strconv.Atoi(timeStringValue)

	if err != nil {
		return 0, err
	}

	switch timeType {
	case "s":
		return timeValue, nil
	case "m":
		return timeValue * 60, nil
	case "h":
		return timeValue * 3600, nil
	}

	return 0, errors.New("Invalid input")
}

func main() {
	var input string

	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		fmt.Println("Please enter the time to wait like: 2s, 3m or 1h:")
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		input = strings.Replace(input, "\r\n", "", -1)
		input = strings.Replace(input, "\n", "", -1)
	}

	seconds, err := getArgsSeconds(input)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	displayDuration(seconds)
	ticker := time.NewTicker(time.Second)
	for {
		_ = <-ticker.C
		seconds--
		displayDuration(seconds)
		if seconds <= 0 {
			ticker.Stop()
			Play()
			break
		}
	}
}
