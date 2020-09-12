package main

import (
	"fmt"
	"strconv"
	"time"
)

func (d data) playGame(limit int) {
	fmt.Println("Quiz starts! Your Questions are:")
	s := "Problem #"
	s2 := ": "
	var score int

	ansCh := make(chan string)

	for i, row := range d {
		qnNo := s + strconv.Itoa(i+1) + s2
		fmt.Print(qnNo, row[0], " ")

		t := time.Tick(time.Duration(limit) * time.Second)

		go getAnswer(ansCh)
		select {
		case ans := <-ansCh:
			if ans == row[1] {
				score = score + 1
			}
		case <-t:
			fmt.Println("\n time over!")
			fmt.Println("you scored", score, "out of", len(d), ".")
			return
		}
	}
	fmt.Println("you scored", score, "out of", len(d), ".")
}

func getAnswer(c chan string) {
	var ans string
	fmt.Scanln(&ans)
	c <- ans
}
