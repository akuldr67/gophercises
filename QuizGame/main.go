package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type data [][]string

func main() {
	csvFlag := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limitFlag := flag.Int("limit", 30, "time limit for the quiz in seconds")
	flag.Parse()

	d := data(readCSV(*csvFlag))
	d.playGame(*limitFlag)
}

func readCSV(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	r := csv.NewReader(f)
	d, error := r.ReadAll()
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	return d
}
