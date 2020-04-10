package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type problem struct {
	question string
	solution int
}

type problems []problem

var numOfQns int
var score int

func main() {
	// Flags
	filenamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeoutPtr := flag.Int("limit", 30, "number of secs before quiz is timed out")

	flag.Parse()

	// Exit if csv file not found
	if fileExists(*filenamePtr) == false {
		fmt.Println(*filenamePtr + " not found!")
		os.Exit(1)
	}

	quiz := readFileByLine(*filenamePtr)

	numOfQns = len(quiz)

	var i int

	go timeout(*timeoutPtr)

	for _, entry := range quiz {
		fmt.Printf("%s : ", entry.question)
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println("Invalid Response that is not of type Int")
		}
		if i == entry.solution {
			score++
		}
	}
	fmt.Printf("Score: %d of %d.\n", score, numOfQns)

}

func readFileByLine(pathToFile string) problems {
	var p problems
	inFile, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println(err.Error() + `: ` + pathToFile)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		tuple := strings.Split(scanner.Text(), ",")
		if v, err := strconv.Atoi(tuple[1]); err == nil {
			p = append(p, problem{tuple[0], v})
		}
	}

	return p
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func timeout(t int) {
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("")
	fmt.Printf("Score: %d of %d.\n", score, numOfQns)
	os.Exit(0)
}
